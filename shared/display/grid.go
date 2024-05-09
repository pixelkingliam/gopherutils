package display

import (
	"errors"
	"github.com/adnsv/go-markout/wcwidth"
	"gopherutils/shared/gquery"
	"strings"
)

func DrawBoxGrid(gridData []string, gridWidth int) (string, error) {

	if gridWidth <= 3 {
		return "", errors.New("int smaller then 3")
	}
	var finalString string
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

	finalString += "┏"
	for i := 0; i < columns; i++ {
		for j := 0; j < wcwidth.StringCells(bigString); j++ {
			finalString += "━"
		}

		if i < columns-1 {
			finalString += "┳"
		}
	}

	finalString += "┓\n"
	// Draw entries
	i := 0
	for x := 0; x < rows; x++ {
		finalString += "┃"
		for y := 0; y < columns; y++ {
			if i >= len(finals) {
				finalString += strings.Repeat(" ", wcwidth.StringCells(bigString)) + "┃" // Fill empty cells with spaces
				continue
			}
			finalString += finals[i]
			i++
			finalString += "┃"
		}
		finalString += "\n"
	}
	// Bottom line
	finalString += "┗"
	for i := 0; i < columns; i++ {
		for j := 0; j < wcwidth.StringCells(bigString); j++ {
			finalString += "━"
		}
		if i < columns-1 {
			finalString += "┻"
		}
	}
	finalString += "┛"
	return finalString, nil
}

func DrawTabGrid(gridData []string, gridWidth int) (string, error) {
	if gridWidth <= 0 {
		return "", errors.New("int smaller then 3")
	}
	var finalString string
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
				finalString += strings.Repeat(" ", wcwidth.StringCells(bigString)) + " " // Fill empty cells with spaces
				continue
			}
			finalString += finals[i]
			finalString += " "
			i++
		}
		finalString += "\n"
	}
	return finalString, nil

}
