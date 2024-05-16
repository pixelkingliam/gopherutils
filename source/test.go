package main

import (
	"fmt"
	"gopherutils/shared/ansi"
	"gopherutils/shared/display"
	"gopherutils/shared/gquery"
	"slices"
)

func main() {

	fmt.Println("[ TESTING ] => gquery")
	var gqueryResults []bool

	// Testing Reverse
	result := testReverse()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> gquery.Reverse[T comparable](slice []T) []T")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> gquery.Reverse[T comparable](slice []T) []T")
	}
	// Testing RemoveAt
	result = testRemoveAt()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> gquery.RemoveAt[T comparable](slice []T, index int) []T")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> gquery.RemoveAt[T comparable](slice []T, index int) []T")
	}
	gqueryResults = append(gqueryResults, result)
	// Testing Swap
	result = testSwap()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> gquery.Swap[T any](slice []T, a int, b int) []T")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> gquery.Swap[T any](slice []T, a int, b int) []T")
	}
	gqueryResults = append(gqueryResults, result)
	// Testing QuickSortLen
	result = testQuick()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> gquery.QuickSortLen(slice []string, pivot int) []string")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> gquery.QuickSortLen(slice []string, pivot int) []string")
	}
	gqueryResults = append(gqueryResults, result)
	// Testing All
	result = testAll()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> gquery.All[T comparable](slice []T, value T) int")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> gquery.All[T comparable](slice []T, value T) int")
	}
	gqueryResults = append(gqueryResults, result)
	// Testing EndsWith
	result = testEndsWith()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> gquery.EndsWith(str string, compare string) bool")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> gquery.EndsWith(str string, compare string) bool")
	}
	gqueryResults = append(gqueryResults, result)
	// Testing StartsWith
	result = testStartsWith()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> gquery.StartsWith(str string, compare string) bool")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> gquery.StartsWith(str string, compare string) bool")
	}
	gqueryResults = append(gqueryResults, result)
	failures := gquery.All(gqueryResults, false)
	if failures == 0 {
		fmt.Println(ansi.GreenFG+"[ ALL TESTS PASS ]"+ansi.ResetColor, "=> gquery")
	} else {
		fmt.Println(ansi.RedFG+"[", failures, "TESTS FAILED]"+ansi.ResetColor, "=> gquery")

	}
	fmt.Println("[ TESTING ] => display")
	var displayResults []bool

	// Testing DynamicBoxGrid
	result = testDynamicBoxGrid()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> display.DynamicBoxGrid(gridData []string, gridWidth int) (string, error)")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> display.DynamicBoxGrid(gridData []string, gridWidth int) (string, error)")
	}
	displayResults = append(displayResults, result)
	// Testing DynamicTabGrid
	result = testDynamicTabGrid()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> display.DynamicTabGrid(gridData []string, gridWidth int) (string, error)")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> display.DynamicTabGrid(gridData []string, gridWidth int) (string, error)")
	}
	displayResults = append(displayResults, result)
	// Testing StaticBoxGrid
	result = testStaticBoxGrid()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> display.StaticBoxGrid(gridData [][]string, options ...bool) (string, error)")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> display.StaticBoxGrid(gridData [][]string, options ...bool) (string, error)")
	}
	displayResults = append(displayResults, result)
	// Testing StaticTabGrid
	result = testStaticBoxGrid()
	if result {
		fmt.Println(ansi.GreenFG+"[ TEST PASSED ]"+ansi.ResetColor, "=> display.StaticTabGrid(gridData [][]string) (string, error)")
	} else {
		fmt.Println(ansi.RedFG+"[ TEST FAILED ]"+ansi.ResetColor, "=> display.StaticTabGrid(gridData [][]string) (string, error)")
	}
	displayResults = append(displayResults, result)

	failures = gquery.All(displayResults, false)
	if failures == 0 {
		fmt.Println(ansi.GreenFG+"[ ALL TESTS PASS ]"+ansi.ResetColor, "=> display")
	} else {
		fmt.Println(ansi.RedFG+"[", failures, "TESTS FAILED]"+ansi.ResetColor, "=> display")
	}

}
func testReverse() bool {
	mySlice := []string{"laughter", "hi", "hell", "process", "x", "one", "water cup", "bottle", "fiver"}
	return slices.Equal(gquery.Reverse(mySlice), []string{"fiver", "bottle", "water cup", "one", "x", "process", "hell", "hi", "laughter"})
}
func testRemoveAt() bool {
	mySlice := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	otherSlice := gquery.RemoveAt(mySlice, 1)
	return slices.Equal(otherSlice, []string{"1", "3", "4", "5", "6", "7", "8", "9"})
}
func testSwap() bool {
	mySlice := gquery.Swap([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}, 0, 2)
	return slices.Equal(mySlice, []string{"3", "2", "1", "4", "5", "6", "7", "8", "9"})

}
func testQuick() bool {
	mySlice := []string{"laughter", "hi", "hell", "process", "x", "one", "wateracup", "bottle", "fiver"}
	mySlice = gquery.QuickSortLen(mySlice)
	expected := []string{"x", "hi", "one", "hell", "fiver", "bottle", "process", "laughter", "wateracup"}
	return slices.Equal(mySlice, expected)
}
func testAll() bool {
	mySlice := []string{"hi", "hi", "hell", "process", "x", "one", "wateracup", "bottle", "fiver"}
	result := gquery.All(mySlice, "hi")
	return result == 2
}
func testEndsWith() bool {
	return gquery.EndsWith("word_letter_number", "number")
}
func testStartsWith() bool {
	return gquery.StartsWith("word_letter_number", "word")
}
func testDynamicBoxGrid() bool {
	slice1d := []string{
		"apple",
		"banana",
		"cherry",
		"date",
		"elderberry",
		"fig",
		"grape",
		"honeydew",
		"kiwi",
		"lemon",
		"mango",
		"nectarine",
		"orange",
		"pear",
		"quince",
		"raspberry",
		"strawberry",
		"tangerine",
		"uva",
		"watermelon",
		"xylocarp",
		"yuzu",
		"zucchini",
		"apricot",
		"blueberry",
		"cranberry",
		"durian",
		"eggplant",
		"feijoa",
		"guava",
		"huckleberry",
		"jackfruit",
		"kiwifruit",
		"lychee",
		"mulberry",
		"nutmeg"}
	output, _ := display.DynamicBoxGrid(slice1d, 228)
	expected :=
		`┏━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┓
┃apple      ┃banana     ┃cherry     ┃date       ┃elderberry ┃fig        ┃grape      ┃honeydew   ┃kiwi       ┃lemon      ┃mango      ┃nectarine  ┃orange     ┃pear       ┃quince     ┃raspberry  ┃strawberry ┃
┃tangerine  ┃uva        ┃watermelon ┃xylocarp   ┃yuzu       ┃zucchini   ┃apricot    ┃blueberry  ┃cranberry  ┃durian     ┃eggplant   ┃feijoa     ┃guava      ┃huckleberry┃jackfruit  ┃kiwifruit  ┃lychee     ┃
┃mulberry   ┃nutmeg     ┃           ┃           ┃           ┃           ┃           ┃           ┃           ┃           ┃           ┃           ┃           ┃           ┃           ┃           ┃           ┃
┗━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┻━━━━━━━━━━━┛`
	return output == expected
}
func testDynamicTabGrid() bool {
	slice1d := []string{
		"apple",
		"banana",
		"cherry",
		"date",
		"elderberry",
		"fig",
		"grape",
		"honeydew",
		"kiwi",
		"lemon",
		"mango",
		"nectarine",
		"orange",
		"pear",
		"quince",
		"raspberry",
		"strawberry",
		"tangerine",
		"uva",
		"watermelon",
		"xylocarp",
		"yuzu",
		"zucchini",
		"apricot",
		"blueberry",
		"cranberry",
		"durian",
		"eggplant",
		"feijoa",
		"guava",
		"huckleberry",
		"jackfruit",
		"kiwifruit",
		"lychee",
		"mulberry",
		"nutmeg"}
	output, _ := display.DynamicTabGrid(slice1d, 228)
	expected := `apple       banana      cherry      date        elderberry  fig         grape       honeydew    kiwi        lemon       mango       nectarine   orange      pear        quince      raspberry   strawberry  
tangerine   uva         watermelon  xylocarp    yuzu        zucchini    apricot     blueberry   cranberry   durian      eggplant    feijoa      guava       huckleberry jackfruit   kiwifruit   lychee      
mulberry    nutmeg                                                                                                                                                                                          
`
	return output == expected
}
func testStaticBoxGrid() bool {
	slice2d := [][]string{
		{"Name", "Age", "Email", "Phone Number", "Address"},
		{"John Doe", "35", "john.doe@example.com", "+1234567890", "123 Main St, Any-town, USA"},
		{"Jane Smith", "28", "jane.smith@gmail.com", "+1987654321", "456 Elm St, Another Town, USA"},
		{"Michael Johnson", "42", "michael.johnson@example.org", "+1122334455", "789 Oak St, Somewhere, USA"},
		{"Emily Davis", "22", "emily.davis@example.net", "+9988776655", "101 Pine St, Nowhere, USA"},
		{"David Brown", "50", "david.brown@example.com", "+1122334455", "1234 Maple Ave, Anywhere, USA"},
		{"Sarah Lee", "31", "sarah.lee@example.com", "+5544332211", "567 Cedar St, Everywhere, USA"},
		{"Daniel Kim", "45", "daniel.kim@example.net", "+3322114455", "890 Birch St, Here, USA"},
		{"Olivia Wilson", "29", "olivia.wilson@example.org", "+1122334455", "345 Oak-wood Ln, There, USA"},
		{"James Taylor", "38", "james.taylor@example.com", "+7788994455", "678 Pinterest Dr, Anywhere-ville, USA"},
		{"Emma Martinez", "26", "emma.martinez@example.net", "+3322114455", "910 Cherry St, Everywhere-ville, USA"},
	}
	noHeader :=
		`┏━━━━━━━━━━━━━━━┳━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃Name           ┃Age┃Email                      ┃Phone Number┃Address                              ┃
┃John Doe       ┃35 ┃john.doe@example.com       ┃+1234567890 ┃123 Main St, Any-town, USA           ┃
┃Jane Smith     ┃28 ┃jane.smith@gmail.com       ┃+1987654321 ┃456 Elm St, Another Town, USA        ┃
┃Michael Johnson┃42 ┃michael.johnson@example.org┃+1122334455 ┃789 Oak St, Somewhere, USA           ┃
┃Emily Davis    ┃22 ┃emily.davis@example.net    ┃+9988776655 ┃101 Pine St, Nowhere, USA            ┃
┃David Brown    ┃50 ┃david.brown@example.com    ┃+1122334455 ┃1234 Maple Ave, Anywhere, USA        ┃
┃Sarah Lee      ┃31 ┃sarah.lee@example.com      ┃+5544332211 ┃567 Cedar St, Everywhere, USA        ┃
┃Daniel Kim     ┃45 ┃daniel.kim@example.net     ┃+3322114455 ┃890 Birch St, Here, USA              ┃
┃Olivia Wilson  ┃29 ┃olivia.wilson@example.org  ┃+1122334455 ┃345 Oak-wood Ln, There, USA          ┃
┃James Taylor   ┃38 ┃james.taylor@example.com   ┃+7788994455 ┃678 Pinterest Dr, Anywhere-ville, USA┃
┃Emma Martinez  ┃26 ┃emma.martinez@example.net  ┃+3322114455 ┃910 Cherry St, Everywhere-ville, USA ┃
┗━━━━━━━━━━━━━━━┻━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛`
	output, err := display.StaticBoxGrid(slice2d)
	if err != nil {
		return false
	}
	if output != noHeader {
		return false
	}
	output, err = display.StaticBoxGrid(slice2d, true)
	if err != nil {
		return false
	}
	header := `┏━━━━━━━━━━━━━━━┳━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃Name           ┃Age┃Email                      ┃Phone Number┃Address                              ┃
┣━━━━━━━━━━━━━━━╋━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━╋━━━━━━━━━━━━╋━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃John Doe       ┃35 ┃john.doe@example.com       ┃+1234567890 ┃123 Main St, Any-town, USA           ┃
┃Jane Smith     ┃28 ┃jane.smith@gmail.com       ┃+1987654321 ┃456 Elm St, Another Town, USA        ┃
┃Michael Johnson┃42 ┃michael.johnson@example.org┃+1122334455 ┃789 Oak St, Somewhere, USA           ┃
┃Emily Davis    ┃22 ┃emily.davis@example.net    ┃+9988776655 ┃101 Pine St, Nowhere, USA            ┃
┃David Brown    ┃50 ┃david.brown@example.com    ┃+1122334455 ┃1234 Maple Ave, Anywhere, USA        ┃
┃Sarah Lee      ┃31 ┃sarah.lee@example.com      ┃+5544332211 ┃567 Cedar St, Everywhere, USA        ┃
┃Daniel Kim     ┃45 ┃daniel.kim@example.net     ┃+3322114455 ┃890 Birch St, Here, USA              ┃
┃Olivia Wilson  ┃29 ┃olivia.wilson@example.org  ┃+1122334455 ┃345 Oak-wood Ln, There, USA          ┃
┃James Taylor   ┃38 ┃james.taylor@example.com   ┃+7788994455 ┃678 Pinterest Dr, Anywhere-ville, USA┃
┃Emma Martinez  ┃26 ┃emma.martinez@example.net  ┃+3322114455 ┃910 Cherry St, Everywhere-ville, USA ┃
┗━━━━━━━━━━━━━━━┻━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━┻━━━━━━━━━━━━┻━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛`
	if output != header {
		return false
	}
	return true
}
func testStaticTabGrid() bool {
	slice2d := [][]string{
		{"Name", "Age", "Email", "Phone Number", "Address"},
		{"John Doe", "35", "john.doe@example.com", "+1234567890", "123 Main St, Any-town, USA"},
		{"Jane Smith", "28", "jane.smith@gmail.com", "+1987654321", "456 Elm St, Another Town, USA"},
		{"Michael Johnson", "42", "michael.johnson@example.org", "+1122334455", "789 Oak St, Somewhere, USA"},
		{"Emily Davis", "22", "emily.davis@example.net", "+9988776655", "101 Pine St, Nowhere, USA"},
		{"David Brown", "50", "david.brown@example.com", "+1122334455", "1234 Maple Ave, Anywhere, USA"},
		{"Sarah Lee", "31", "sarah.lee@example.com", "+5544332211", "567 Cedar St, Everywhere, USA"},
		{"Daniel Kim", "45", "daniel.kim@example.net", "+3322114455", "890 Birch St, Here, USA"},
		{"Olivia Wilson", "29", "olivia.wilson@example.org", "+1122334455", "345 Oak-wood Ln, There, USA"},
		{"James Taylor", "38", "james.taylor@example.com", "+7788994455", "678 Pinterest Dr, Anywhere-ville, USA"},
		{"Emma Martinez", "26", "emma.martinez@example.net", "+3322114455", "910 Cherry St, Everywhere-ville, USA"},
	}
	expected :=
		` Name            Age Email                       Phone Number Address                               
 John Doe        35  john.doe@example.com        +1234567890  123 Main St, Any-town, USA            
 Jane Smith      28  jane.smith@gmail.com        +1987654321  456 Elm St, Another Town, USA         
 Michael Johnson 42  michael.johnson@example.org +1122334455  789 Oak St, Somewhere, USA            
 Emily Davis     22  emily.davis@example.net     +9988776655  101 Pine St, Nowhere, USA             
 David Brown     50  david.brown@example.com     +1122334455  1234 Maple Ave, Anywhere, USA         
 Sarah Lee       31  sarah.lee@example.com       +5544332211  567 Cedar St, Everywhere, USA         
 Daniel Kim      45  daniel.kim@example.net      +3322114455  890 Birch St, Here, USA               
 Olivia Wilson   29  olivia.wilson@example.org   +1122334455  345 Oak-wood Ln, There, USA           
 James Taylor    38  james.taylor@example.com    +7788994455  678 Pinterest Dr, Anywhere-ville, USA 
 Emma Martinez   26  emma.martinez@example.net   +3322114455  910 Cherry St, Everywhere-ville, USA `
	output, err := display.StaticTabGrid(slice2d)
	if err != nil {
		return false
	}
	if output != expected {
		return false
	}

	return true
}
