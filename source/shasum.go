package main

import (
	"bytes"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {
	var options struct {
		Algorithm int  `short:"a" long:"algorithm" description:"Selects the SHA algorithm to use, available options:1 (Default)\n224\n256\n384\n512\n512224\n512256"`                   // GNU Compatible
		Binary    bool `short:"b" long:"binary" description:"Reads in binary mode, does nothing on GNU systems."`                                                                       // GNU Compatible
		Text      bool `short:"t" long:"text" description:"Reads in text mode."`                                                                                                        // GNU Compatible
		Tag       bool `short:"T" long:"tag" description:"Writes BSD-style checksums."`                                                                                                 // GNU Compatible
		BitsMode  bool `short:"0" long:"01" description:"Reads in BITS mode.\nASCII '0' is interpreted as 0-bit\nASCII '1' is interpreted as 1-bit\nAll other characters are ignored."` // GNU Compatible
		Universal bool `short:"U" long:"UNIVERSAL" description:"Reads in Universal newlines mode.\n\tNormalizes different newline formats to LF ('\n')"`                                // GNU Compatible

	}
	options.Text = true
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

	if !checkAlgo(options.Algorithm) {
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
	if false {
		// TODO CHECK ARG
	} else {
		for i := 0; i < len(args); i++ {
			file, err := os.ReadFile(args[i])
			if options.BitsMode {
				file = readBitsMode(file)
			}
			if options.Universal {
				file = bytes.Replace(file, []byte{'\r', '\n'}, []byte{'\n'}, -1)
				file = bytes.Replace(file, []byte{'\r'}, []byte{'\n'}, -1)
			}
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			fmt.Printf("%s%s", formatHash(options.Algorithm, file, options.Tag, args[i], mode), "\n")
		}
	}

}
func readBitsMode(data []byte) []byte {
	var length = 0
	for _, b := range data {
		if b != '1' && b != '0' {
			continue
		}
		length++
	}
	var bytes = make([]byte, (length+7)/8)

	var tByte = uint8(0)
	var iByte = 0
	var iBit = 0
	for _, b := range data {
		if b != '1' && b != '0' {
			continue
		}
		if b == '1' {
			tByte |= 1 << (7 - iBit)
		}
		iBit++

		if iBit == 8 {
			iBit = 0
			bytes[iByte] = tByte
			tByte = uint8(0)
			iByte++
		}
	}
	if iBit != 0 {
		bytes[iByte] = tByte
	}
	for _, b := range bytes {
		fmt.Printf("%08b\n", b)
	}
	return bytes
}
func checkAlgo(algorithm int) bool {
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
func algoString(algorithm int) string {
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
func getHash(algorithm int, data []byte) (string, error) {
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
func formatHash(algorithm int, data []byte, tag bool, fileName string, mode int) string {
	hash, err := getHash(algorithm, data)
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
		return fmt.Sprintf("%s (%s) = %s", algoString(algorithm), fileName, hash)
	} else {
		return fmt.Sprintf("%s %s%s", hash, indicator, fileName)
	}
}
