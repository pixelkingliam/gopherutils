package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"golang.org/x/term"
	"gopherutils/shared/ansi"
	"gopherutils/shared/gquery"
	"os"
	"slices"
)

func main() {
	//goland:noinspection ALL
	var options struct {
		NoColor   bool `short:"c" long:"nocolors" description:"Render results with colors"`
		List      bool `short:"l" long:"long" description:"Use long list."`
		SortByLen bool `short:"s" long:"sortlen" description:"QuickSortLen using string length instead of alphabetical sorting."`
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
		starting = ansi.BlueFG
		ending += ansi.ResetColor
	}

	for i := 0; i < len(dirs); i++ {
		dirs[i] = starting + dirs[i] + ending
	}
	if !term.IsTerminal(int(os.Stdin.Fd())) || options.List {
		if options.SortByLen {
			entries := slices.Concat(gquery.Reverse(gquery.QuickSortLen(dirs)), gquery.Reverse(gquery.QuickSortLen(files)))
			for i := 0; i < len(entries); i++ {
				fmt.Println(entries[i])
			}
		} else {
			entries := slices.Concat(dirs, files)
			for i := 0; i < len(entries); i++ {
				fmt.Println(entries[i])
			}
		}

		os.Exit(0)
	}

}
