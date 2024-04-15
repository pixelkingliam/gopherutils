package main

import (
	"fmt"
	gquery "gopherutils/shared"
	"slices"
)

func main() {

	fmt.Println("[ TESTING ] => gquery")
	var gqueryResults []bool
	const green = "\033[38;5;2m"
	const red = "\033[38;5;1m"
	const clearColor = "\033[0m"

	// Testing Reverse
	result := testReverse()
	if result {
		fmt.Println(green+"[ TEST PASSED ]"+clearColor, "=> gquery.Reverse[T comparable](slice []T) []T")
	} else {
		fmt.Println(red+"[ TEST FAILED ]"+clearColor, "=> gquery.Reverse[T comparable](slice []T) []T")
	}
	// Testing RemoveAt
	result = testRemoveAt()
	if result {
		fmt.Println(green+"[ TEST PASSED ]"+clearColor, "=> gquery.RemoveAt[T comparable](slice []T, index int) []T")
	} else {
		fmt.Println(red+"[ TEST FAILED ]"+clearColor, "=> gquery.RemoveAt[T comparable](slice []T, index int) []T")
	}
	gqueryResults = append(gqueryResults, result)
	// Testing Swap
	result = testSwap()
	if result {
		fmt.Println(green+"[ TEST PASSED ]"+clearColor, "=> gquery.Swap[T any](slice []T, a int, b int) []T")
	} else {
		fmt.Println(red+"[ TEST FAILED ]"+clearColor, "=> gquery.Swap[T any](slice []T, a int, b int) []T")
	}
	gqueryResults = append(gqueryResults, result)
	// Testing QuickSortLen
	result = testQuick()
	if result {
		fmt.Println(green+"[ TEST PASSED ]"+clearColor, "=> gquery.QuickSortLen(slice []string, pivot int) []string")
	} else {
		fmt.Println(red+"[ TEST FAILED ]"+clearColor, "=> gquery.QuickSortLen(slice []string, pivot int) []string")
	}
	gqueryResults = append(gqueryResults, result)
	// Testing All
	result = testAll()
	if result {
		fmt.Println(green+"[ TEST PASSED ]"+clearColor, "=> gquery.All[T comparable](slice []T, value T) int")
	} else {
		fmt.Println(red+"[ TEST FAILED ]"+clearColor, "=> gquery.All[T comparable](slice []T, value T) int")
	}
	gqueryResults = append(gqueryResults, result)
	failures := gquery.All(gqueryResults, false)
	if failures == 0 {
		fmt.Println(green+"[ ALL TESTS PASS ]"+clearColor, "=> gquery")
	} else {
		fmt.Println(red+"[", failures, "TESTS FAILED]"+clearColor, "=> gquery")

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
