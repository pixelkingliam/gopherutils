package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"gopherutils/shared/hashing"
	"io"
	"os"
)

func main() {
	var options struct {
		Algorithm     int  `short:"a" long:"algorithm" description:"Selects the SHA algorithm to use, available options:1 (Default)\n224\n256\n384\n512\n512224\n512256"`                   // GNU Compatible
		Binary        bool `short:"b" long:"binary" description:"Reads in binary mode, does nothing on GNU systems."`                                                                       // GNU Compatible
		Text          bool `short:"t" long:"text" description:"Reads in text mode."`                                                                                                        // GNU Compatible
		Tag           bool `short:"T" long:"tag" description:"Writes BSD-style checksums."`                                                                                                 // GNU Compatible
		BitsMode      bool `short:"0" long:"01" description:"Reads in BITS mode.\nASCII '0' is interpreted as 0-bit\nASCII '1' is interpreted as 1-bit\nAll other characters are ignored."` // GNU Compatible
		Universal     bool `short:"U" long:"UNIVERSAL" description:"Reads in Universal newlines mode.\n\tNormalizes different newline formats to LF ('\n')"`                                // GNU Compatible
		Check         bool `short:"c" long:"check" description:"Reads checksums from FILEs and verifies them."`                                                                             // GNU Compatible
		Warn          bool `short:"w" long:"warn" description:"Writes a warning for each mal-formated line."`                                                                               // GNU Compatible
		Status        bool `short:"s" long:"status" description:"Avoids printing, rely on exit status code instead."`                                                                       // GNU Compatible
		Quiet         bool `short:"q" long:"quiet" description:"Avoids printing \"OK\" for each successfully verified file."`                                                               // GNU Compatible
		IgnoreMissing bool `short:"i" long:"ignore-missing" description:"Ignores missing files instead of fail"`                                                                            // GNU Compatible
		Strict        bool `short:"S" long:"strict" description:"Exit non-zero for improperly formatted checksum lines."`                                                                   // GNU Compatible
	}
	options.Algorithm = 1
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
	if !shaCheckAlgo(options.Algorithm) {
		fmt.Println("Invalid SHA algorithm\nTry 'shasum -h' for help.")
		os.Exit(1)
	}
	readStdIn := false
	if len(args) == 0 {
		readStdIn = true
	}
	if args[0] == "-" {
		readStdIn = true
	}
	if readStdIn {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Printf("Error reading stdin: %s\n", err.Error())
			os.Exit(1)
		}
		var sum hashing.Sum
		if options.BitsMode {
			sum.Mode |= hashing.BitMode
		}
		if options.Universal {
			sum.Mode |= hashing.Universal
		}
		if options.Tag {
			sum.Mode |= hashing.Tag
		}
		if options.Binary {
			sum.Mode |= hashing.Binary
		}

		sum.HashType = shaAlgoString(options.Algorithm)
		sum.File = "-"
		sum.Hash = hashing.Hash(data[:], sum.HashType)
		fmt.Println(sum)
		os.Exit(0)
	}

	if options.Check {
		failed := 0
		malFormatted := 0
		notExist := 0
		if options.Tag {
			fmt.Printf("--tag option is incompatible with --check.\n")
			return
		}
		sums := make([]hashing.Sum, 0)
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
			readSums, malFormattedLines := hashing.ReadSums(string(data), "sha")
			malFormatted += malFormattedLines
			sums = append(sums, readSums...)
		}
		for _, sum := range sums {
			verifySum, err := hashing.VerifySum(sum)
			if err != nil {
				if err.Error() == "file does not exist" {
					notExist++
				} else if err.Error() == "invalid sum algorithm" {
					fmt.Println("Internal error")
					os.Exit(1)
				}
			}
			if verifySum {
				if !options.Status && !options.Quiet {
					fmt.Printf("%s: OK\n", sum.File)
				}
			} else {
				failed++
				if !options.Status {
					fmt.Printf("%s: FAILED\n", sum.File)
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
	} else {
		for i := 0; i < len(args); i++ {
			var template hashing.SumTemplate
			if options.BitsMode {
				template.Mode |= hashing.BitMode

			}
			if options.Universal {
				template.Mode |= hashing.Universal

			}
			if options.Tag {
				template.Mode |= hashing.Tag
			}
			if options.Binary {
				template.Mode |= hashing.Binary
			}
			template.HashType = shaAlgoString(options.Algorithm)
			template.File = args[i]
			sum := hashing.GetSum(template)
			fmt.Println(sum)
		}
	}

}
func shaCheckAlgo(algorithm int) bool {
	switch algorithm {
	case 1:
		return true
	case 224:
		return true
	case 256:
		return true
	case 384:
		return true
	case 512:
		return true
	case 512224:
		return true
	case 512256:
		return true
	default:
		return false

	}
}
func shaAlgoString(algorithm int) string {
	switch algorithm {
	case 1:
		return "SHA1"
	case 224:
		return "SHA224"
	case 256:
		return "SHA256"
	case 384:
		return "SHA384"
	case 512:
		return "SHA512"
	case 512224:
		return "SHA512/224"
	case 512256:
		return "SHA512/256"
	default:
		return "UNKNOWN"

	}
}
