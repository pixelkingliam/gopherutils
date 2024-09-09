package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"syscall"
	"time"
)

func main() {
	var options struct {
		NoCreate         bool `short:"c" long:"no-create" description:"Does not create any files."` // GNU Compatible
		AccessOnly       bool `short:"a"  description:"Only changes the access time"`               // GNU Compatible
		ModificationOnly bool `short:"m"  description:"Only changes the modification time"`         // GNU Compatible
		//VArg
	}
	args, err := flags.ParseArgs(&options, os.Args)

	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	//VCode
	if len(args) == 1 {
		fmt.Println("Missing file operand")
		os.Exit(1)
	}
	args = args[1:]

	for _, arg := range args {
		stat, err := os.Stat(arg)

		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				if options.NoCreate {
					continue
				}
				_, err2 := os.Create(arg)
				if err2 != nil {
					fmt.Printf("Unknown error: %s\n", err.Error())
					os.Exit(1)
				}

			} else {
				fmt.Printf("Unknown error: %s\n", err.Error())
				os.Exit(1)
			}
		}

		mtime := time.Now()
		atime := time.Now()

		if options.AccessOnly {
			mtime = stat.ModTime()

		}
		if options.ModificationOnly {
			sysAtime := stat.Sys().(*syscall.Stat_t).Atim
			atime = time.Unix(sysAtime.Sec, sysAtime.Nsec)
		}
		err = os.Chtimes(arg, atime, mtime)
		if err != nil {
			fmt.Printf("Unknown error: %s\n", err.Error())
			os.Exit(1)
		}

	}

}
