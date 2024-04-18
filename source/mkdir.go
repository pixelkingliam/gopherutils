package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var options struct {
		Parents bool `short:"p" long:"parents" description:"No errors if existing, also creates necessary parent directories as needed."`
	}

	args, err := flags.ParseArgs(&options, os.Args)
	if err != nil {
		fmt.Println(`Unexpected error:`, err.Error())
	}
	if len(args) == 0 {
		fmt.Println("Missing operand.")
		os.Exit(1)
	}
	dir := args[0]
	if options.Parents {
		err = os.MkdirAll(dir, 0755)
	} else {
		err = os.Mkdir(dir, 0755)
	}
	if err != nil {
		fmt.Println(`Unexpected error:`, err.Error())
		os.Exit(1)
	}
}
