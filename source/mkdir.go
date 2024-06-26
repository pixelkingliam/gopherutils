package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var options struct {
		Parents bool `short:"p" long:"parents" description:"No errors if existing, also creates necessary parent directories as needed."` // GNU Compatible
		Mode    int  `short:"m" long:"mode" description:"Set file mode (numerical representation)" default:"0755"`                        // GNU Compatible
		Verbose bool `short:"v" long:"verbose" description:"Print a message for every directories created"`                               // GNU Compatible
	}
	options.Mode = 0755
	args, err := flags.ParseArgs(&options, os.Args)
	args = args[1:]
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	if len(args) == 0 {
		fmt.Println("Missing operand.")
		os.Exit(1)
	}
	for i := 0; i < len(args); i++ {
		dir := args[i]
		if options.Parents {
			err = os.MkdirAll(dir, os.FileMode(options.Mode))
		} else {
			err = os.Mkdir(dir, os.FileMode(options.Mode))
		}
		if err != nil {
			fmt.Println(`Unexpected error:`, err.Error())
			os.Exit(1)
		} else if options.Verbose {
			fmt.Printf("Creating directory %s.\n", dir)
		}

	}
}
