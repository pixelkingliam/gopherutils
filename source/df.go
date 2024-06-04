package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/shirou/gopsutil/disk"
	"gopherutils/shared/convert"
	"gopherutils/shared/display"
	"gopherutils/shared/gquery"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var options struct {
		BinaryPrefix bool   `short:"b" long:"human-readable" description:"Displays sizes in powers of 1024 (e.g., 1023MiB)"`
		SIPrefix     bool   `short:"H" long:"si" description:"print sizes in powers of 1000 (e.g., 1.1GB)"`
		TabTable     bool   `short:"T" long:"tab" description:"Displays the table using tabs; GNU Compatible"` // By default, renders with UTF-8 grid
		Posix        bool   `short:"P" long:"portability" description:"Uses POSIX-compatible header; implies -P"`
		All          bool   `short:"a" long:"all" description:"Include all file systems"`
		BlockSize    string `short:"B" long:"block-size" description:"Scale sizes by SIZE before printing them. The SIZE parameter specifies the units in which sizes should be displayed.\nFor example, '-BM' prints sizes in units of 1,048,576 bytes (1MiB), '-BK' prints sizes in units of 1,024 bytes (1KiB), and so on.\nSupported suffixes for SIZE include: 'B' (bytes),\n'K' (Kilobytes, 1024 bytes),\n'M' (Megabytes, 1024^2 bytes),\n'G' (Gigabytes, 1024^3 bytes),\n'T' (Terabytes, 1024^4 bytes),\n'P' (Petabytes, 1024^5 bytes),\nand 'E' (Exabytes, 1024^6 bytes)\n"`
		INodes       bool   `short:"i" long:"inodes" description:"List inode information instead of block usage"`
	}
	options.BlockSize = "1K"
	_, err := flags.ParseArgs(&options, os.Args[1:])
	options.BlockSize = sanitizeBArg(options.BlockSize)
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	if options.Posix {
		options.TabTable = true
		if options.BlockSize == "1K" {
			options.BlockSize = "1024"
		}
	}
	partitions, err := disk.Partitions(options.All)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	table := make([][]string, 0)
	if options.INodes {
		table = append(table, []string{"Device", "Type", "Inodes", "IUsed", "IFree", "IUse%", "Mounted On"})

	} else {
		table = append(table, []string{"Device", "Type", fmt.Sprintf("%s-blocks", options.BlockSize), "Used", "Available", "Use%", "Mounted On"})

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
		if options.INodes {
			table = append(table, []string{partition.Device, partition.Fstype, fmt.Sprintf("%v", usage.InodesTotal), fmt.Sprintf("%v", usage.InodesUsed), fmt.Sprintf("%v", usage.InodesFree), fmt.Sprintf("%v%%", int(math.Round(usage.InodesUsedPercent))), partition.Mountpoint})

			continue
		}
		size := fmt.Sprintf("%v", resolveBlocks(int64(usage.Total), options.BlockSize))
		if options.SIPrefix {
			size = convert.ToSI(usage.Total, true)
		} else if options.BinaryPrefix {
			size = convert.ToBinary(usage.Total, true)

		}

		used := fmt.Sprintf("%v", resolveBlocks(int64(usage.Used), options.BlockSize))
		if options.SIPrefix {
			used = convert.ToSI(usage.Used, true)
		} else if options.BinaryPrefix {
			used = convert.ToBinary(usage.Used, true)

		}

		free := fmt.Sprintf("%v", resolveBlocks(int64(usage.Total), options.BlockSize)-resolveBlocks(int64(usage.Used), options.BlockSize))
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
func resolveBlocks(bytes int64, blocks string) int64 {
	var v int64
	if gquery.IsDigit(blocks[len(blocks)-1]) {
		v, _ = strconv.ParseInt(blocks, 10, 64)
		return bytes / v
	}
	v, _ = strconv.ParseInt(blocks[:len(blocks)-1], 10, 64)
	unitChar := strings.ToLower(blocks)[len(blocks)-1]
	var unitModifier int64
	switch unitChar {
	case 'k':
		unitModifier = 1024
	case 'm':
		unitModifier = 1024 * 1024
	case 'g':
		unitModifier = 1024 * 1024 * 1024
	case 't':
		unitModifier = 1024 * 1024 * 1024 * 1024
	case 'p':
		unitModifier = 1024 * 1024 * 1024 * 1024 * 1024
	case 'e':
		unitModifier = 1024 * 1024 * 1024 * 1024 * 1024 * 1024
	default:
		unitModifier = 1 // bytes
	}
	return bytes / (v * unitModifier)
}

func sanitizeBArg(str string) string {
	var builder strings.Builder
	for i := 0; i < len(str); i++ {
		if !gquery.IsDigit(str[i]) {
			builder.WriteString(strings.ToLower(string(str[i])))
			break
		}
		builder.WriteByte(str[i])
	}
	final := builder.String()
	if len(final) > 0 {
		lastChar := final[len(final)-1]
		if unicode.IsLetter(rune(lastChar)) {
			switch lastChar {
			case 'k', 'm', 'g', 't', 'p', 'e':
				return final
			default:
				return final[:len(final)-1]
			}
		} else {
			return final[:len(final)-1]
		}
	}

	return final
}
