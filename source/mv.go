package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"io"
	"os"
)

func main() {
	var options struct {
		NoOverwrite    bool `short:"n" long:"no-clobber" description:"Prevents overwriting existing files."` // GNU Compatible
		ForceOverwrite bool `short:"f" long:"force" description:"Do not prompt before overwriting files."`   // GNU Compatible
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
	if len(args) == 0 {
		fmt.Println("Missing file operand.")
		os.Exit(1)
	}
	if len(args) == 1 {
		fmt.Println("Missing destination file operand after '" + args[0] + "'")
		os.Exit(1)
	}
	if len(args) == 2 {
		src, err := os.Open(args[0])
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Printf("Cannot open `%s`: No such file or directory\n", args[0])
			} else {
				fmt.Println(err.Error())
			}

			os.Exit(1)
		}
		dest, err := os.Create(args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		_, err = io.Copy(dest, src)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		stat, err := os.Stat(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = os.Chmod(args[1], stat.Mode())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		err = os.Remove(args[0])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		os.Exit(0)

	}

	// TODO Multiple moves
	dest := args[len(args)-1]
	fmt.Println(dest)

}
