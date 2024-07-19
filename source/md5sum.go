package main

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"strings"
)

func main() {
	var options struct {
		Zero   bool `short:"z" long:"zero" description:"Ends each output with a NUL character instead of a newline character."` // GNU Compatible
		Binary bool `short:"b" long:"binary" description:"Reads in binary mode, does nothing on GNU systems."`                  // GNU Compatible
		Text   bool `short:"t" long:"text" description:"Reads in text mode."`                                                   // GNU Compatible
		Tag    bool `long:"tag" description:"Writes BSD-style checksums."`                                                      // GNU Compatible
		Check  bool `short:"c" long:"check" description:"Reads checksums from FILEs and verifies them."`                        // GNU Compatible
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
	if options.Check {
		failed := 0
		if options.Tag {
			fmt.Printf("--tag option is incompatible with --check.\n")
			return
		}
		for _, file := range args {
			_, err = os.Stat(file)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					fmt.Printf("%s: No such file.\n", file)
				}

				return
			}
			data, err := os.ReadFile(file)
			if err != nil {
				fmt.Printf("Unknown error! %s", err)
				return
			}
			lines := strings.Split(string(data), "\n")
			for _, line := range lines {
				if line == "" {
					continue
				}
				hash := strings.Split(line, " ")[0]
				var hashFilePath string
				if strings.Split(line, " ")[1] == "" {
					hashFilePath = strings.Split(line, " ")[2]
				} else {
					hashFilePath = strings.Split(line, " ")[1][1:]
				}
				_, err = os.Stat(hashFilePath)
				if err != nil {
					fmt.Printf("Unknown error! %s", err)
					return
				}
				hashFile, err := os.ReadFile(hashFilePath)
				if err != nil {
					fmt.Printf("Unknown error! %s", err)
					return
				}
				if hash == fmt.Sprintf("%x", md5.Sum(hashFile)) {
					fmt.Printf("%s: OK\n", hashFilePath)
				} else {
					failed++
					fmt.Printf("%s: FAILED\n", hashFilePath)
				}
			}
		}
		if failed != 0 {
			fmt.Printf("WARNING: %v checksum did NOT match\n", failed)
			os.Exit(1)
		}
		os.Exit(0)
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
