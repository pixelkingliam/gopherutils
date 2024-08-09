package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"golang.org/x/crypto/blake2b"
	"gopherutils/shared/convert"
	"os"
	"strings"
)

func main() {
	var options struct {
		Length        int  `short:"l" long:"length" description:"The length of the hash in bits, must be a multiple of 8 and mustn't exceed the multiple of the blake2 algorithm"`          // GNU Compatible
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
	options.Text = true
	options.Length = 512
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

	if !b2CheckAlgo(options.Length) {
		fmt.Println("Invalid BLAKE2b length\nTry 'b2sum -h' for help.")
		os.Exit(1)
	}
	mode := 0
	if options.Binary {
		mode = 1
	}
	if options.BitsMode {
		mode = 2
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
				options.Length = b2AlgoFromLength(len(strings.Split(line, " ")[0]))
				hashLength := b2LengthAlgo(options.Length)
				hash := line[:hashLength]
				if len(line) < hashLength+3 || !strings.Contains(line, " ") || len(strings.Split(line, " ")[0]) != hashLength {
					if options.Warn && !options.Status {
						fmt.Printf("Line %v is improperly formatted.\n", atLine+1)
					}
					malFormatted++
					continue
				}
				hashFilePath := line[hashLength+2:]
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
				calculated, err := b2GetHash(options.Length, hashFile)
				if err != nil {
					fmt.Printf("Unknown error calculating hash.")
					os.Exit(1)
				}
				if hash == calculated {
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
	} else {
		for i := 0; i < len(args); i++ {
			file, err := os.ReadFile(args[i])
			if options.BitsMode {
				file = convert.ReadAsciiBits(file)
			}
			if options.Universal {
				file = bytes.Replace(file, []byte{'\r', '\n'}, []byte{'\n'}, -1)
				file = bytes.Replace(file, []byte{'\r'}, []byte{'\n'}, -1)
			}
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			fmt.Printf("%s%s", b2FormatHash(options.Length, file, options.Tag, args[i], mode), "\n")
		}
	}

}
func b2AlgoFromLength(length int) int {
	return length * 4
}
func b2CheckAlgo(algorithm int) bool {
	return algorithm%8 == 0
}
func b2LengthAlgo(algorithm int) int {
	return algorithm / 4
}

func b2AlgoString(algorithm int) string {
	return fmt.Sprintf("BLAKE2b-%d", algorithm)
}
func b2GetHash(algorithm int, data []byte) (string, error) {
	hash, err := blake2b.New(algorithm/8, nil)
	if err != nil {
		return "", err
	}
	hash.Write(data)
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
func b2FormatHash(algorithm int, data []byte, tag bool, fileName string, mode int) string {
	hash, err := b2GetHash(algorithm, data)
	indicator := " "
	if mode == 1 {
		indicator = "*"
	}
	if err != nil {
		fmt.Printf("Error getting hash for %s: %s\n", fileName, err)
		os.Exit(1)
	}
	if tag {
		return fmt.Sprintf("%s (%s) = %s", b2AlgoString(algorithm), fileName, hash)
	} else {
		return fmt.Sprintf("%s %s%s", hash, indicator, fileName)
	}
}
