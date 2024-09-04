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
		Zero         bool `short:"z" long:"zero" description:"Ends each output line with NUL, instead of newline."`                            // GNU Incompatible
		CanonMissing bool `short:"m" long:"canonicalize-missing" description:"Suppresses error messages associated with missing directories."` // GNU Incompatible
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
	terminator := "\n"
	if options.Zero {
		terminator = "\x00"
	}
	for _, fakePath := range args {
		final := make([]string, 0)
		if fakePath[0] != '/' {
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
				return
			}
			final = strings.Split(pwd, "/")
		}
		for i, str := range strings.Split(fakePath, "/") {
			if str == ".." || str == "../" {
				if i == 0 {
					continue
				}
				_, err := os.Stat(formPath(final))
				if err != nil && !options.CanonMissing {
					fmt.Printf("Error: %v\n", err)
					return
				}
				final = final[:len(final)-1]
				continue
			}
			path := append(final, str)
			lstat, err := os.Lstat(formPath(final))
			if err != nil {
				final = path
				continue
			}
			if lstat.Mode()&os.ModeSymlink != 0 {
				target, err := os.Readlink(formPath(path))
				if err != nil {
					fmt.Printf("Error getting symlink location: %v\n", err)
					return
				}
				final = strings.Split(target, "/")
			} else {
				final = path
			}
		}

		fmt.Printf("%s%s", formPath(final), terminator)
	}
}
func formPath(components []string) string {
	sorted := make([]string, 0)
	for _, str := range components {
		if str != "" {
			sorted = append(sorted, str)
		}
	}
	components = sorted
	result := "/"
	for _, component := range components {
		result += component
		result += "/"
	}
	return result[:len(result)-1]
}
