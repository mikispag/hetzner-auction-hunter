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