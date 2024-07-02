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
		Directory       bool `short:"d" long:"parents" description:"No errors if existing, also creates necessary parent directories as needed."`                             // GNU Compatible
		Number          bool `short:"n" long:"number" description:"Numbers all output lines."`                                                                                // GNU Compatible
		NumberNonBlank  bool `short:"b" long:"number-nonblank" description:"Numbers all non-blank output lines."`                                                             // GNU Compatible
		OmitBlank       bool `short:"o" long:"omit-blank" description:"Avoids printing blank lines."`                                                                         // GNU Compatible Addition
		ShowEnds        bool `short:"E" long:"show-ends" description:"Display $ at the end of each line."`                                                                    // GNU Compatible
		ShowTabs        bool `short:"T" long:"show-tabs" description:"Displays TAB characters as ^I."`                                                                        // GNU Compatible
		Ignored         bool `short:"u" long:"ignored" description:"Ignored."`                                                                                                // GNU Compatible
		SqueezeBlank    bool `short:"s" long:"squeeze-blank" description:"Avoids printing repeated blank lines"`                                                              // GNU Compatible
		ShowNonPrinting bool `short:"v" long:"show-nonprinting" description:"Prints control characters and meta characters using ^ and M- notation, except for LFD and TAB."` // GNU Compatible
		VE              bool `short:"e" description:"Equivalent to -VE"`                                                                                                      // GNU Compatible
		VT              bool `short:"t" description:"Equivalent to -VT"`                                                                                                      // GNU Compatible
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
	if options.VE {
		options.ShowNonPrinting = true
		options.ShowEnds = true
	}
	if options.VT {
		options.ShowNonPrinting = true
		options.ShowTabs = true
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
		if options.ShowNonPrinting {
			file = []byte(processNonPrinting(string(file)))
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
		if i != len(lines) {
			fmt.Printf("\n")
		}
	}
}
func processNonPrinting(str string) string {
	var output strings.Builder
	for i := 0; i < len(str); i++ {
		char := str[i]
		if char == 127 {
			output.WriteByte('^')
			output.WriteByte('?')
			continue
		}
		if char <= 31 {
			// print TAB and Newline, otherwise represent control character with ^ notation
			if char != '\t' && char != '\n' {
				output.WriteByte('^')
				output.WriteByte(char + 64)
				continue
			}
		}
		if char == 255 {
			output.WriteByte('M')
			output.WriteByte('-')
			output.WriteByte('?')
			continue
		}
		if char > 127 {
			output.WriteByte('M')
			output.WriteByte('-')
			// convert meta character into ASCII character, then plus +64 to assure printable
			output.WriteByte(char - 128 + 64)
			continue
		}
		output.WriteByte(char)

	}
	return output.String()
}
