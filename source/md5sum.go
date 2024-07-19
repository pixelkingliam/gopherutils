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
		Warn   bool `short:"w" long:"warn" description:"Writes a warning for each mal-formated line."`                          // GNU Compatible
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
	if !options.Check {
		if options.Warn {
			fmt.Printf("The --warn/-w option depends on --check.\n see --help for more information.\n")
			os.Exit(1)
		}
	}
	if options.Check {
		failed := 0
		malformat := 0
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
			for atLine, line := range lines {
				if line == "" {
					continue
				}
				hash := line[:32]
				if len(line) < 35 || !strings.Contains(line, " ") || len(strings.Split(line, " ")[0]) != 32 {
					if options.Warn {
						fmt.Printf("Line %v is improperly formatted.\n", atLine+1)
					}
					malformat++
					continue
				}
				hashFilePath := line[34:]
				//hashFilePath = hashFilePath[1:]
				/*var hashFilePath string
				if strings.Split(line, " ")[1] == "" {
					hashFilePath = strings.Split(line, " ")[2]
				} else {
					hashFilePath = strings.Split(line, " ")[1][1:]
				}*/
				_, err = os.Stat(hashFilePath)
				if err != nil {
					if errors.Is(err, os.ErrNotExist) {
						fmt.Printf("File '%s' does not exist", hashFilePath)
						os.Exit(1)
					}
					fmt.Printf("Unknown error! %s", err)
					os.Exit(1)
				}
				hashFile, err := os.ReadFile(hashFilePath)
				if err != nil {
					fmt.Printf("Unknown error reading file! %s", err)
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
		exit := 0
		if failed != 0 {
			fmt.Printf("WARNING: %v checksum did NOT match\n", failed)
			exit = 1
		}
		if malformat != 0 {
			fmt.Printf("WARNING: %v line is improperly formatted.\n", malformat)
			exit = 1
		}
		os.Exit(exit)
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
