package main

import (
	"fmt"
	"github.com/adnsv/go-markout/wcwidth"
	"github.com/jessevdk/go-flags"
	"golang.org/x/term"
	"gopherutils/shared/ansi"
	"gopherutils/shared/gquery"

	"os"
	"slices"
	"strings"
)

func main() {
	//goland:noinspection ALL
	var options struct {
		NoColor   bool `short:"c" long:"nocolors" description:"Render results with ansi"`
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
	ending := ""
	if !options.NoColor {
		starting = ansi.BlueFG
		ending += ansi.ResetColor
	}
	if !term.IsTerminal(int(os.Stdin.Fd())) || options.List {

		for i := 0; i < len(dirs); i++ {
			dirs[i] = starting + dirs[i] + "/" + ending
		}
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
	} else {
		//entries := slices.Concat(gquery.Reverse(gquery.QuickSortLen(dirs)), gquery.Reverse(gquery.QuickSortLen(files)))
		totalLen := len(dirs) + len(files)
		width, _, err := term.GetSize(int(os.Stdin.Fd()))
		if err != nil {
			return
		}
		bigString := dirs[gquery.BiggestString(dirs)]
		if wcwidth.StringCells(bigString) < wcwidth.StringCells(files[gquery.BiggestString(files)]) {
			bigString = files[gquery.BiggestString(files)]
		}

		columns := (width) / wcwidth.StringCells(bigString)
		rows := (totalLen + columns - 1) / columns
		if rows == 1 && columns > totalLen {
			columns = totalLen
		}

		// Add padding to all entries
		var finals []string
		for i := 0; i < len(dirs); i++ {
			val := dirs[i]

			val = val + "/"
			val = val + strings.Repeat(" ", wcwidth.StringCells(bigString)-wcwidth.StringCells(val))
			// starting and ending here are ANSI colorization
			finals = append(finals, starting+val+ending)
		}
		for i := 0; i < len(files); i++ {
			val := files[i]

			val = val + strings.Repeat(" ", wcwidth.StringCells(bigString)-wcwidth.StringCells(val))
			finals = append(finals, val)
		}
		//for _, final := range finals {
		//
		//	fmt.Println(final, "|")
		//	//for _, runeValue := range final {
		//	//	fmt.Printf("%U '%c'\n", runeValue, runeValue)
		//	//}
		//}
		fmt.Println("BIG STRING IS", bigString)
		// Top line

		fmt.Print("┏")
		for i := 0; i < columns; i++ {
			for j := 0; j < wcwidth.StringCells(bigString); j++ {
				fmt.Print("━")
			}

			if i < columns-1 {
				fmt.Print("┳")
			}
		}

		fmt.Print("┓\n")
		// Draw entries
		i := 0
		for x := 0; x < rows; x++ {
			fmt.Print("┃")
			for y := 0; y < columns; y++ {
				if i >= len(finals) {
					fmt.Print(strings.Repeat(" ", wcwidth.StringCells(bigString)) + "┃") // Fill empty cells with spaces
					continue
				}
				fmt.Print(finals[i])
				i++
				fmt.Print("┃")
			}
			fmt.Println()
		}
		// Bottom line
		fmt.Print("┗")
		for i := 0; i < columns; i++ {
			for j := 0; j < wcwidth.StringCells(bigString); j++ {
				fmt.Print("━")
			}
			if i < columns-1 {
				fmt.Print("┻")
			}
		}
		fmt.Print("┛\n")
		fmt.Println(rows, columns)
		fmt.Println(totalLen, width)
	}

}

//for _, entry := range entries {
//	stringLen := len(ansi.StripANSI(entry))
//
//	padding := bigStringLen - stringLen
//	if padding < 0 {
//		padding = 0
//	}
//
//	finals = append(finals, entry+strings.Repeat(" ", padding))
//}
