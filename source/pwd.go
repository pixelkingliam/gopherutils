package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var options struct {
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
	result, _ := os.Getwd()
	fmt.Println(result)
}
