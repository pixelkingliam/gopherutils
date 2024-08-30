package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var options struct {
		Zero bool `short:"z" long:"zero" description:"Ends each output line with NUL, instead of newline."` // GNU Incompatible
	}
	args, err := flags.ParseArgs(&options, os.Args)
	if len(args) != 0 {
		args = args[1:]

	}
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	}
	terminator := "\n"
	if options.Zero {
		terminator = "\x00"
	}
	for _, fakepath := range args {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}
		fmt.Printf("%s/%s%s", pwd, fakepath, terminator)
	}
}
