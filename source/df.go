package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/shirou/gopsutil/disk"
	"gopherutils/shared/convert"
	"gopherutils/shared/display"
	"math"
	"os"
)

func main() {
	var options struct {
		BinaryPrefix bool `short:"b" long:"human-readable" description:"Displays sizes in powers of 1024 (e.g., 1023MiB)"`
		SIPrefix     bool `short:"H" long:"si-prefix" description:"print sizes in powers of 1000 (e.g., 1.1GB)"`
		TabTable     bool `short:"T" long:"tab" description:"Displays the table using tabs; GNU Compatible"`
		Posix        bool `short:"P" long:"portability" description:"Uses POSIX-compatible header; implies -P"`
		All          bool `short:"a" long:"all" description:"Include all file systems"`
	}
	_, err := flags.ParseArgs(&options, os.Args[1:])
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	if options.Posix {
		options.TabTable = true
	}
	partitions, err := disk.Partitions(options.All)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	table := make([][]string, 0)
	if options.Posix {
		table = append(table, []string{"Device", "Type", "1024-blocks", "Used", "Available", "Capacity", "Mounted On"})
		options.TabTable = true
	} else {
		table = append(table, []string{"Device", "Type", "1K-blocks", "Used", "Available", "Use%", "Mounted On"})
	}
	if options.BinaryPrefix || options.SIPrefix {
		table[0][2] = "Size"
	}
	for _, partition := range partitions {
		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		size := fmt.Sprintf("%v", usage.Total/1024)
		if options.SIPrefix {
			size = convert.ToSI(usage.Total, true)
		} else if options.BinaryPrefix {
			size = convert.ToBinary(usage.Total, true)

		}

		used := fmt.Sprintf("%v", usage.Used/1024)
		if options.SIPrefix {
			used = convert.ToSI(usage.Used, true)
		} else if options.BinaryPrefix {
			used = convert.ToBinary(usage.Used, true)

		}

		free := fmt.Sprintf("%v", usage.Total/1024-usage.Used/1024)
		if options.SIPrefix {
			free = convert.ToSI(usage.Total-usage.Used, true)
		} else if options.BinaryPrefix {
			free = convert.ToBinary(usage.Total-usage.Used, true)

		}
		table = append(table, []string{partition.Device, partition.Fstype, size, used, free, fmt.Sprintf("%v%%", int(math.Round(usage.UsedPercent))), partition.Mountpoint})
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
