package main

import (
	"errors"
	"fmt"
	"github.com/jessevdk/go-flags"
	"math/rand"
	"os"
	"strings"
)

func main() {
	var options struct {
		Directory bool `short:"d" long:"parents" description:"No errors if existing, also creates necessary parent directories as needed."` // GNU Compatible
		DryRun    bool `short:"u" long:"dry-run" description:"Do not create anything. Only prints a name."`
	}

	args, err := flags.ParseArgs(&options, os.Args)
	args = args[1:]
	if err != nil {
		if errors.Is(err, flags.ErrHelp) {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	if len(args) > 1 {
		fmt.Println("Too many templates")
		os.Exit(1)
	}
	template := "/tmp/tmp.XXXXXXXXXX"
	if len(args) == 1 {
		template = args[0]
	}

	path := generatePath(template)
	for exists(path) {
		path = generatePath(template)
	}
	if !options.DryRun {
		if options.Directory {
			err := os.Mkdir(path, os.FileMode(0755))
			if err != nil {
				fmt.Printf("Unexpected error: %s\n", err.Error())
				os.Exit(1)
			}
		} else {
			_, err := os.Create(path)
			if err != nil {
				fmt.Printf("Unexpected error: %s\n", err.Error())
				os.Exit(1)
			}
		}
	}
	fmt.Println(path)
}
func generatePath(template string) string {
	var builder strings.Builder
	charSet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < len(template); i++ {
		if template[i] == 'X' {
			randomChar := charSet[rand.Intn(len(charSet))]
			builder.WriteByte(randomChar)
		} else {
			builder.WriteByte(template[i])
		}
	}
	return builder.String()
}
func exists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
