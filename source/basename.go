package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"gopherutils/shared/gquery"
	"os"
	"path/filepath"
)

func main() {
	//goland:noinspection ALL
	var options struct {
		Multiple bool   `short:"a" long:"multiple" description:"Enabled by default; For GNU-Compatibility"`             // GNU Compatible
		Zero     bool   `short:"z" long:"zero" description:"End each output line with a NULL byte instead of newline."` // GNU Compatible
		Suffix   string `short:"s" long:"suffix" description:"Removes trailing SUFFIX in all outputs."`                 // GNU Compatible
	}

	args, err := flags.ParseArgs(&options, os.Args)
	args = args[1:]
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	if options.Multiple {
		os.Stderr.WriteString("Warning: -a/--multiple is deprecated and is now enabled by default.\n")
	}
	ending := "\n"
	if options.Zero {
		ending = "\x00"
	}
	for _, arg := range args {
		if gquery.EndsWith(arg, options.Suffix) {
			arg = arg[:len(arg)-len(options.Suffix)]
		}
		fmt.Printf(filepath.Base(arg) + ending)
	}

}
