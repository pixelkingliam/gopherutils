package main

import (
	"errors"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var options struct {
		//VArg
	}
	//args
	_, err := flags.ParseArgs(&options, os.Args[1:])
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	//VCode
	os.Exit(0)
}
