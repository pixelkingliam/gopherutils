package main

import (
	"fmt"
	"github.com/shirou/gopsutil/disk"
	"os"
)

func main() {
	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s (%s) at %s with %s (%f) \n", partition.Device, disk.GetDiskSerialNumber(partition.Device), partition.Mountpoint, partition.Opts, usage.UsedPercent)
	}
}
