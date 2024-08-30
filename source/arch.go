package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"runtime"
)

func main() {
	var options struct {
		Language bool `short:"l" long:"language" description:"Display the architecture name as defined by the implementation language instead of the standard GNU-compatible architecture name."` // GNU Compatible

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

	output := runtime.GOARCH
	if !options.Language {
		switch runtime.GOARCH {
		case "amd64":
			output = "x86_64"
			break
		case "386":
			output = "i386"
		case "arm":
			output = "armv8l"
		case "arm64":
			output = "aarch64"
		case "ppc64":
			output = "ppc64"
		case "ppc64le":
			output = "ppc64le"
		case "s390x":
			output = "s390x"
		case "risc64":
			output = "riscv64"
		case "loong64":
			output = "loong64"
		case "mips":
			output = "mips"
		case "mipsle":
			output = "mips"
		case "mips64":
			output = "mips64"
		case "mips64le":
			output = "mips64"
		}
	}
	println(output)
}
