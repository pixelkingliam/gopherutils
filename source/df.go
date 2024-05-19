package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/shirou/gopsutil/disk"
	"gopherutils/shared/display"
	"math"
	"os"
)

func main() {
	var options struct {
		TabTable bool `short:"T" long:"tab" description:"Displays the table using tabs; GNU Compatible"`
	}
	args, err := flags.ParseArgs(&options, os.Args)
	args = args[1:]
	partitions, err := disk.Partitions(false)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	table := make([][]string, 0)
	table = append(table, []string{"Device", "Type", "1K-blocks", "Used", "Available", "Use%", "Mounted On"})
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		table = append(table, []string{partition.Device, partition.Fstype, fmt.Sprintf("%v", usage.Total/1024), fmt.Sprintf("%v", usage.Used/1024), fmt.Sprintf("%v", usage.Total/1024-usage.Used/1024), fmt.Sprintf("%v%%", int(math.Round(usage.UsedPercent))), partition.Mountpoint})
	}
	var output string
	if options.TabTable {
		output, err = display.StaticTabGrid(table)

	} else {
		output, err = display.StaticBoxGrid(table, true)

	}
	if err != nil {
		fmt.Println("Render error:", err.Error())
	}
	fmt.Println(output)
}
