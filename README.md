# Hetzner server auction hunter

[![Go Report](https://goreportcard.com/badge/github.com/mikispag/hetzner-auction-hunter)](https://goreportcard.com/badge/github.com/mikispag/hetzner-auction-hunter)


Print Hetzner server auctions satisfying custom parameters.

## Usage
```console
Usage of hetzner-auction-hunter:
  -ecc
    	whether to require ECC RAM
  -highio
    	whether to require a fast disk
  -maxPrice float
    	the maximum price per month to pay, in EUR (default 9.223372036854776e+18)
  -minCPUBenchmark uint
    	the minimum CPU benchmark score
  -minCPUCount uint
    	the minimum amount of CPUs
  -minDisk uint
    	the minimum total disk size, in GB
  -minDiskCount uint
    	the minimum amount of disks
  -minRAM uint
    	the minimum RAM size, in GB
```

Example usage:

```console
$ hetzner-auction-hunter -minDiskCount 4 -minDisk 12000 -maxPrice 38
EUR 35.29 / month
	4x RAM 4096 MB DDR3 4x HDD SATA 3,0 TB RAID Controller 4-Port SATA PCI-E - LSI MegaRAID SAS 9260-4i 4x 3 TB Intel Core i7-2600 SB42 1282573 HWR SAS
	CPU: 1x Intel Core i7-2600 (benchmark: 5343)
	RAM: 16 GB (ECC: false)
	Disks: 4x 3 TB (SSD/High IO: false)
	Datacenter: [FSN1-DC6 FSN]

EUR 37.82 / month
	4x RAM 8192 MB DDR3 4x HDD SATA 3,0 TB RAID Controller 4-Port SATA PCI-E - LSI MegaRAID SAS 9260-4i 4x 3 TB Intel Core i7-3770 SB45 1277367 HWR SAS
	CPU: 1x Intel Core i7-3770 (benchmark: 6370)
	RAM: 32 GB (ECC: false)
	Disks: 4x 3 TB (SSD/High IO: false)
	Datacenter: [FSN1-DC5 FSN]

EUR 37.82 / month
	4x RAM 8192 MB DDR3 4x HDD SATA 3,0 TB 4x 3 TB Intel Core i7-2600 SB45 1282585
	CPU: 1x Intel Core i7-2600 (benchmark: 5343)
	RAM: 32 GB (ECC: false)
	Disks: 4x 3 TB (SSD/High IO: false)
	Datacenter: [FSN1-DC7 FSN]
```