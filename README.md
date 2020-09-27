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
$ hetzner-auction-hunter -minDiskCount 4 -minDisk 16000 -maxPrice 50
EUR 42.86 / month (next reduction on Fri, 02 Oct 2020 00:01:58 CEST)
	2x RAM 8192 MB DDR3 4x HDD SATA 4,0 TB Enterprise 4x 4 TB Intel Core i7-2600 SB51 1283233 Ent. HDD
	CPU: 1x Intel Core i7-2600 (benchmark: 5343)
	RAM: 16 GB (ECC: false)
	Disks: 4x 4 TB (SSD/High IO: false)
	Datacenter: [FSN1-DC11 FSN]

EUR 46.22 / month (next reduction on Tue, 29 Sep 2020 05:32:57 CEST)
	4x RAM 8192 MB DDR3 4x HDD SATA 4,0 TB Enterprise 4x 4 TB Intel Core i7-2600 SB55 1287368 Ent. HDD
	CPU: 1x Intel Core i7-2600 (benchmark: 5343)
	RAM: 32 GB (ECC: false)
	Disks: 4x 4 TB (SSD/High IO: false)
	Datacenter: [FSN1-DC3 FSN]

EUR 47.90 / month (next reduction on Fri, 02 Oct 2020 11:09:59 CEST)
	2x RAM 8192 MB DDR3 4x HDD SATA 4,0 TB Enterprise NIC 1 Gbit - Argus LR-9201 (1x1GbE RJ45) 4x 4 TB Intel Core i7-2600 SB57 1284543 Ent. HDD
	CPU: 1x Intel Core i7-2600 (benchmark: 5343)
	RAM: 16 GB (ECC: false)
	Disks: 4x 4 TB (SSD/High IO: false)
	Datacenter: [FSN1-DC12 FSN]
```