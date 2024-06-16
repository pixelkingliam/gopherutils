package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"time"
)

func main() {
	var options struct {
		NoCreate bool `short:"c" long:"no-create" description:"Does not create any files."` // GNU Compatible

	}
	args, err := flags.ParseArgs(&options, os.Args)

	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	if len(args) == 1 {
		fmt.Println("Missing file operand")
		os.Exit(1)
	}
	args = args[1:]

	for _, arg := range args {
		_, err := os.Stat(arg)
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
		err = os.Chtimes(arg, time.Now(), time.Now())
		if err != nil {
			fmt.Printf("Unknown error: %s\n", err.Error())
			os.Exit(1)
		}

	}

}
