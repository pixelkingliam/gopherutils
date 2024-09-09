package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var options struct {
		Null bool `short:"0" long:"null" description:"Ends each output line with NUL instead of a newline."`
		//VArg
	}
	_, err := flags.ParseArgs(&options, os.Args[1:])
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	//VCode
	ending := '\n'
	for _, str := range os.Environ() {
		if str[:2] != "_=" {
			fmt.Printf("%s%c", str, ending)

		}
	}
}
