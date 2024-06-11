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
	path := generatePath()
	for exists(path) {
		path = generatePath()
	}
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
	fmt.Println(path)
}
func generatePath() string {
	var builder strings.Builder
	charSet := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	for i := 0; i < 10; i++ {
		randomChar := charSet[rand.Intn(len(charSet))]
		builder.WriteByte(randomChar)
	}
	return "/tmp/tmp." + builder.String()
}
func exists(path string) bool {
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}
