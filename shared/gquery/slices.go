package gquery

import (
	"github.com/adnsv/go-markout/wcwidth"
	"slices"
)

func Remove[T comparable](slice []T, element T) []T {
	return RemoveAt(slice, slices.Index(slice, element))
}
func RemoveAt[T comparable](slice []T, index int) []T {
	if index > len(slice) {
		return slice
	}
	result := make([]T, 0)
	result = append(result, slice[:index]...)
	result = append(result, slice[index+1:]...)
	return result

}
func Reverse[T comparable](slice []T) []T {
	var output []T
	for i := len(slice) - 1; i >= 0; i-- {
		output = append(output, slice[i])
	}
	return output
}
func BiggestString(slice []string) int {
	biggestIndex := 0

	for i := 0; i < len(slice); i++ {
		if wcwidth.StringCells(slice[biggestIndex]) < wcwidth.StringCells(slice[i]) {
			biggestIndex = i
		}
	}
	return biggestIndex

}
func Swap[T any](slice []T, a int, b int) []T {
	if len(slice) <= a || len(slice) <= b || a < 0 || b < 0 || b == a {
		return slice
	}
	temp := slice[a]
	slice[a] = slice[b]
	slice[b] = temp
	return slice
}

func QuickSortLen(slice []string) []string {
	if len(slice) <= 1 {
		return slice
	}
	pivot := len(slice) / 2
	pivotLength := len(slice[pivot])
	var left []string
	var right []string

	// Partition
	for i, value := range slice {
		if i != pivot {
			if len(value) < pivotLength {
				left = append(left, value)
			} else {
				right = append(right, value)
			}
		}
	}
	//    return    =    QuickSortLen(left) +    Pivot    + QuickSortLen(right)
	return append(append(QuickSortLen(left), slice[pivot]), QuickSortLen(right)...)
}
func All[T comparable](slice []T, value T) int {
	count := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] == value {
			count++
		}
	}
	return count
}
