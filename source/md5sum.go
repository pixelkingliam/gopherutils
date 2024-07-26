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
		Zero          bool `short:"z" long:"zero" description:"Ends each output with a NUL character instead of a newline character."` // GNU Compatible
		Binary        bool `short:"b" long:"binary" description:"Reads in binary mode, does nothing on GNU systems."`                  // GNU Compatible
		Text          bool `short:"t" long:"text" description:"Reads in text mode."`                                                   // GNU Compatible
		Tag           bool `long:"tag" description:"Writes BSD-style checksums."`                                                      // GNU Compatible
		Check         bool `short:"c" long:"check" description:"Reads checksums from FILEs and verifies them."`                        // GNU Compatible
		Warn          bool `short:"w" long:"warn" description:"Writes a warning for each mal-formated line."`                          // GNU Compatible
		Status        bool `short:"s" long:"status" description:"Avoids printing, rely on exit status code instead."`                  // GNU Compatible
		Quiet         bool `short:"q" long:"quiet" description:"Avoids printing \"OK\" for each successfully verified file."`          // GNU Compatible
		IgnoreMissing bool `short:"i" long:"ignore-missing" description:"Ignores missing files instead of fail"`                       // GNU Compatible
		Strict        bool `short:"S" long:"strict" description:"Exit non-zero for improperly formatted checksum lines."`              // GNU Compatible
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
		if options.Status {
			fmt.Printf("The --status/-s option depends on --check.\n see --help for more information.\n")
			os.Exit(1)
		}
	}
	if options.Check {
		failed := 0
		malFormatted := 0
		notExist := 0
		if options.Tag {
			fmt.Printf("--tag option is incompatible with --check.\n")
			return
		}
		for _, file := range args {
			data, err := os.ReadFile(file)
			if err != nil {
				if errors.Is(err, os.ErrNotExist) {
					fmt.Printf("File '%s' does not exist", file)
					os.Exit(1)
				}

				fmt.Printf("Unknown error! %s", err)
				os.Exit(1)
			}
			lines := strings.Split(string(data), "\n")
			for atLine, line := range lines {
				if line == "" {
					continue
				}
				hash := line[:32]
				if len(line) < 35 || !strings.Contains(line, " ") || len(strings.Split(line, " ")[0]) != 32 {
					if options.Warn && !options.Status {
						fmt.Printf("Line %v is improperly formatted.\n", atLine+1)
					}
					malFormatted++
					continue
				}
				hashFilePath := line[34:]

				_, err = os.Stat(hashFilePath)
				if err != nil {

					if !options.IgnoreMissing {
						fmt.Printf("%s: FAILED open or read\n", hashFilePath)
						notExist++
					}
					continue

				}
				hashFile, err := os.ReadFile(hashFilePath)
				if err != nil {
					fmt.Printf("Unknown error reading file! %s", err)
					os.Exit(1)
				}
				if hash == fmt.Sprintf("%x", md5.Sum(hashFile)) {
					if !options.Status && !options.Quiet {
						fmt.Printf("%s: OK\n", hashFilePath)
					}
				} else {
					failed++
					if !options.Status {
						fmt.Printf("%s: FAILED\n", hashFilePath)
					}
				}
			}
		}
		exit := 0
		if failed != 0 && !options.Status {
			fmt.Printf("WARNING: %v checksum did NOT match\n", failed)
			exit = 1
		}
		if malFormatted != 0 && !options.Status {
			fmt.Printf("WARNING: %v line is improperly formatted.\n", malFormatted)
			if options.Strict {
				exit = 1
			}
		}
		if notExist != 0 && !options.Status {
			fmt.Printf("WARNING: %v listed file could not bread.\n", notExist)
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
