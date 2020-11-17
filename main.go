package main

import (
	"flag"
	"fmt"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/mikispag/utils/web"
	log "github.com/sirupsen/logrus"
)

const (
	hetznerAPIURL = "https://www.hetzner.com/a_hz_serverboerse/live_data.json"
)

var (
	c = &http.Client{Timeout: 10 * time.Second}
)

func main() {
	maxPrice := flag.Float64("maxPrice", math.MaxInt64, "the maximum price per month to pay, in EUR")
	minCPUCount := flag.Uint("minCPUCount", 0, "the minimum amount of CPUs")
	minCPUBenchmark := flag.Uint("minCPUBenchmark", 0, "the minimum CPU benchmark score")
	minRAMSize := flag.Uint("minRAM", 0, "the minimum RAM size, in GB")
	minDiskCount := flag.Uint("minDiskCount", 0, "the minimum amount of disks")
	minDiskSize := flag.Uint("minDisk", 0, "the minimum total disk size, in GB")
	ecc := flag.Bool("ecc", false, "whether to require ECC RAM")
	highio := flag.Bool("highio", false, "whether to require a fast disk")
	flag.Parse()

	// Initialize logger
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC822,
	})

	var result map[string]interface{}
	err := web.GetJSON(hetznerAPIURL, &result)
	if err != nil {
		log.WithError(err).Fatal("Unable to fetch servers from Hetzner!")
	}

	for _, server := range result["server"].([]interface{}) {
		s := server.(map[string]interface{})
		price, err := strconv.ParseFloat(s["price"].(string), 64)
		if err != nil {
			log.WithError(err).Fatal("Unable to parse price!")
		}
		if price > *maxPrice {
			continue
		}
		var nextReduce string
		if s["fixed_price"].(bool) {
			nextReduce = "fixed"
		} else {
			nextReduceSeconds := time.Duration(s["next_reduce"].(float64)) * time.Second
			reductionTime := time.Now().Add(nextReduceSeconds)
			nextReduce = "next reduction on " + reductionTime.Format(time.RFC1123)
		}

		cpuCount := uint(s["cpu_count"].(float64))
		if cpuCount < *minCPUCount {
			continue
		}
		cpuBenchmark := uint(s["cpu_benchmark"].(float64))
		if cpuBenchmark < *minCPUBenchmark {
			continue
		}
		ram := uint(s["ram"].(float64))
		if ram < *minRAMSize {
			continue
		}
		if *ecc && !s["is_ecc"].(bool) {
			continue
		}
		diskCount := uint(s["hdd_count"].(float64))
		if diskCount < *minDiskCount {
			continue
		}
		disk := uint(s["hdd_size"].(float64))
		if disk*diskCount < *minDiskSize {
			continue
		}
		if *highio && !s["is_highio"].(bool) {
			continue
		}
		fmt.Printf("EUR %.2f / month (%s)\n\t%s\n", price, nextReduce, s["freetext"].(string))
		fmt.Printf("\tCPU: %dx %s (benchmark: %d)\n", cpuCount, s["cpu"].(string), cpuBenchmark)
		fmt.Printf("\tRAM: %s (ECC: %t)\n", s["ram_hr"].(string), s["is_ecc"].(bool))
		fmt.Printf("\tDisks: %s (SSD/High IO: %t)\n", s["hdd_hr"].(string), s["is_highio"].(bool))
		fmt.Printf("\tDatacenter: %v\n\n", s["datacenter"].([]interface{}))
	}
}
