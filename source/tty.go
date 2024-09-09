package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"golang.org/x/term"
	"os"
)

func ttyname(fd uintptr) (string, error) {
	path := fmt.Sprintf("/proc/self/fd/%v", fd) // Linux-specific
	tty, err := os.Readlink(path)
	if err != nil {
		return "", err
	}
	return tty, nil
}
func main() {
	var options struct {
		Silent bool `short:"s" long:"silent" description:"print nothing, only return an exit status"`
		Quiet  bool `short:"q" long:"quiet" description:"same as -s"`
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
	if options.Quiet {
		options.Silent = true
	}
	if term.IsTerminal(int(os.Stdin.Fd())) {
		result, err := ttyname(os.Stdin.Fd())
		if err != nil {
			if !options.Silent {
				fmt.Println("Unexpect error!", err.Error())
			}
			os.Exit(0)
		}
		if !options.Silent {
			fmt.Println(result)
		}
	} else {
		if !options.Silent {
			fmt.Println("Standard input is not a terminal.")
		}
		os.Exit(1)
	}

}
