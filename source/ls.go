package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"golang.org/x/term"
	"os"
	"slices"
)

func main() {
	//goland:noinspection ALL
	var options struct {
		NoColor bool `short:"c" long:"nocolors" description:"Render results with colors"`
		List    bool `short:"l" long:"long" description:"Use long list."`
	}

	dir := "."
	args, err := flags.ParseArgs(&options, os.Args)

	if err != nil {
		println(err.Error())
		os.Exit(1)
	}

	if len(args) > 1 {
		dir = args[1]
	}
	result, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf(err.Error())
		os.Exit(1)
	}

	var dirs []string
	var files []string
	for i := 0; i < len(result); i++ {
		if result[i].IsDir() {
			dirs = append(dirs, result[i].Name())
		} else {
			files = append(files, result[i].Name())
		}
	}
	slices.Sort(dirs)
	slices.Sort(files)
	starting := ""
	ending := "/"
	if !options.NoColor {
		starting = "\033[38;5;75m"
		ending += "\033[0m"
	}

	for i := 0; i < len(dirs); i++ {
		dirs[i] = starting + dirs[i] + ending
	}

	if !options.NoColor {
		dirs[0] = "\033[38;5;75m" + dirs[0]
		dirs[len(dirs)-1] = dirs[len(dirs)-1] + "\033[0m"
	}

	entries := slices.Concat(dirs, files)
	if !term.IsTerminal(int(os.Stdin.Fd())) || options.List {
		for i := 0; i < len(entries); i++ {
			fmt.Println(entries[i])
		}
		os.Exit(0)
	}

}
