package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"os/user"
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
	res, err := user.Current()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(res.Name)
}
