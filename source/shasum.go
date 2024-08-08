package main

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"gopherutils/shared/convert"
	"os"
	"strings"
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
	autoDetect := true
	options.Text = true
	options.Algorithm = -1
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
	if options.Algorithm != -1 {
		autoDetect = false
	} else {
		options.Algorithm = 1
	}
	if !shaCheckAlgo(options.Algorithm) {
		fmt.Println("Invalid SHA algorithm\nTry 'shasum -h' for help.")
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
				if autoDetect {
					options.Algorithm = shaAlgoFromLength(len(strings.Split(line, " ")[0]))
				}
				hashLength := shaLengthAlgo(options.Algorithm)
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
				calculated, err := shaGetHash(options.Algorithm, hashFile)
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

			fmt.Printf("%s%s", shaFormatHash(options.Algorithm, file, options.Tag, args[i], mode), "\n")
		}
	}

}
func shaLengthAlgo(algorithm int) int {
	switch algorithm {
	case 1:
		return 40
	case 224:
		return 56
	case 256:
		return 64
	case 384:
		return 96
	case 512:
		return 128
	case 512224:
		return 56
	case 512256:
		return 64
	default:
		return -1

	}
}
func shaAlgoFromLength(length int) int {
	switch length {
	case 40:
		return 1
	case 56:
		return 224 // or 512224
	case 64:
		return 256 // or 512256
	case 96:
		return 384
	case 128:
		return 512
	default:
		return -1
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
func shaGetHash(algorithm int, data []byte) (string, error) {
	switch algorithm {
	case 1:
		return fmt.Sprintf("%x", sha1.Sum(data)), nil
	case 224:
		return fmt.Sprintf("%x", sha256.Sum224(data)), nil
	case 256:
		return fmt.Sprintf("%x", sha256.Sum256(data)), nil
	case 384:
		return fmt.Sprintf("%x", sha512.Sum384(data)), nil
	case 512:
		return fmt.Sprintf("%x", sha512.Sum512(data)), nil
	case 512224:
		return fmt.Sprintf("%x", sha512.Sum512_224(data)), nil
	case 512256:
		return fmt.Sprintf("%x", sha512.Sum512_256(data)), nil
	default:
		return "", errors.New("invalid SHA algorithm")

	}
}
func shaFormatHash(algorithm int, data []byte, tag bool, fileName string, mode int) string {
	hash, err := shaGetHash(algorithm, data)
	indicator := " "
	if mode == 1 {
		indicator = "*"
	}
	if mode == 2 {
		indicator = "^"
	}
	if err != nil {
		fmt.Printf("Error getting hash for %s: %s\n", fileName, err)
		os.Exit(1)
	}
	if tag {
		return fmt.Sprintf("%s (%s) = %s", shaAlgoString(algorithm), fileName, hash)
	} else {
		return fmt.Sprintf("%s %s%s", hash, indicator, fileName)
	}
}
