package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"runtime"
	"unicode"
)

func main() {
	var options struct {
		KernelName bool `short:"s" long:"kernel-name" description:"Prints the kernel's name'"` // GNU Compatible
		//VArg
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
	//VCode
	if options.KernelName {
		kernelName := runtime.GOOS
		if kernelName == "windows" {
			kernelName = "NT Kernel"
		} else {
			// Capitalize first letter
			kernelName = fmt.Sprintf("%c%s", unicode.ToUpper(rune(kernelName[0])), kernelName[1:])
		}
		fmt.Print(kernelName)
	}
	fmt.Println()
}
