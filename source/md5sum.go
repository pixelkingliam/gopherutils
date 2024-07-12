package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var options struct {
		Zero   bool `short:"z" long:"zero" description:"Ends each output with a NUL character instead of a newline character."` // GNU Compatible
		Binary bool `short:"b" long:"binary" description:"Reads in binary mode, does nothing on GNU systems."`                  // GNU Compatible
		Text   bool `short:"t" long:"text" description:"Reads in text mode."`                                                   // GNU Compatible
		Tag    bool `long:"tag" description:"Writes BSD-style checksums."`                                                      // GNU Compatible
	}
	options.Text = true
	args, err := flags.ParseArgs(&options, os.Args)
	if len(args) != 0 {
		args = args[1:]

	}
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	ending := "\n"
	if options.Zero {
		ending = "\x00"
	}
	if options.Binary {
		options.Text = false
	}
	prefix := " "
	if !options.Text {
		prefix = "*"
	}
	for _, arg := range args {
		_, err := os.Stat(arg)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Printf("File '%s' does not exist.\n", arg)
			} else {
				fmt.Println(err.Error())
			}
			os.Exit(1)
		}
		file, err := os.ReadFile(arg)
		if err != nil {
			os.Exit(1)
		}
		if options.Tag {
			fmt.Printf("MD5 (%s) = %x%s", arg, md5.Sum(file), ending)
		} else {
			fmt.Printf("%x %s%s%s", md5.Sum(file), prefix, arg, ending)
		}
	}

}
