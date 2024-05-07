package display

import (
	"fmt"
	"github.com/adnsv/go-markout/wcwidth"
	"gopherutils/shared/gquery"
	"strings"
)

func DrawBoxGrid(gridData []string, gridWidth int) {

	bigString := gridData[gquery.BiggestString(gridData)]
	columns := (gridWidth - len(gridData)) / wcwidth.StringCells(bigString)
	rows := (len(gridData) + columns - 1) / columns
	if rows == 1 && columns > len(gridData) {
		columns = len(gridData)
	}
	// Add padding to all entries
	var finals []string
	for i := 0; i < len(gridData); i++ {
		val := gridData[i]

		val = val + strings.Repeat(" ", wcwidth.StringCells(bigString)-wcwidth.StringCells(val))
		finals = append(finals, val)
	}
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

}

func DrawTabGrid(gridData []string, gridWidth int) {
	bigString := gridData[gquery.BiggestString(gridData)]
	columns := (gridWidth - len(gridData)) / wcwidth.StringCells(bigString)
	rows := (len(gridData) + columns - 2) / columns
	if rows == 1 && columns > len(gridData) {
		columns = len(gridData)
	}
	// Add padding to all entries
	var finals []string
	for i := 0; i < len(gridData); i++ {
		val := gridData[i]

		val = val + strings.Repeat(" ", wcwidth.StringCells(bigString)-wcwidth.StringCells(val))
		finals = append(finals, val)
	}

	// Draw entries
	i := 0
	for x := 0; x < rows; x++ {
		for y := 0; y < columns; y++ {
			if i >= len(finals) {
				fmt.Print(strings.Repeat(" ", wcwidth.StringCells(bigString)) + " ") // Fill empty cells with spaces
				continue
			}
			fmt.Print(finals[i])
			fmt.Print(" ")
			i++
		}
		fmt.Println()
	}

}
