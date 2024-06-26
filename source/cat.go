package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
	"strings"
)

func main() {
	var options struct {
		Directory      bool `short:"d" long:"parents" description:"No errors if existing, also creates necessary parent directories as needed."` // GNU Compatible
		Number         bool `short:"n" long:"number" description:"Numbers all output lines."`                                                    // GNU Compatible
		NumberNonBlank bool `short:"b" long:"number-nonblank" description:"Numbers all non-blank output lines."`                                 // GNU Compatible
		OmitBlank      bool `short:"o" long:"omit-blank" description:"Avoids printing blank lines."`                                             // GNU Compatible
		ShowEnds       bool `short:"E" long:"show-ends" description:"Display $ at the end of each line."`                                        // GNU Compatible
		ShowTabs       bool `short:"T" long:"show-tabs" description:"Displays TAB characters as ^I."`                                            // GNU Compatible
		Ignored        bool `short:"u" long:"ignored" description:"Ignored."`                                                                    // GNU Compatible
		SqueezeBlank   bool `short:"s" long:"squeeze-blank" description:"Avoids printing repeated blank lines"`
	}
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
	if options.NumberNonBlank {
		options.Number = true
	}
	var lines []string
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

		if options.ShowTabs {
			lines = append(lines, strings.Split(strings.Replace(string(file), "\t", "^I", -1), "\n")...)
		} else {
			lines = append(lines, strings.Split(string(file), "\n")...)

		}

	}
	lineCount := 1
	for i, line := range lines {
		if len(line) == 0 {
			if options.OmitBlank {
				continue
			}
			if options.SqueezeBlank && i != 0 {
				if len(lines[i-1]) == 0 {
					continue
				}
			}
		}

		if options.Number {
			lineCountStr := fmt.Sprintf("%v", lineCount)
			if options.NumberNonBlank {
				if len(line) == 0 {
					if len(lines)-1 == i {
						os.Exit(0)
					}
					lineCountStr = ""
				}
			}
			fmt.Printf("    %s  %s", lineCountStr, line)
			if options.NumberNonBlank {
				if len(line) != 0 {
					lineCount++
				}
			} else {
				lineCount++

			}
		} else {
			fmt.Printf(line)

		}
		if options.ShowEnds {
			fmt.Printf("$")
		}
		fmt.Printf("\n")
	}
}
