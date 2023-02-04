package main

import (
	_ "fmt"
	"math"
)

// func mainTest() {
// 	a := []int{0, 10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
// 	fmt.Println(StartSearch(100, a))
// }

// Function that takes (tar) as a target value and a slice (s) where tar must be found.
func StartSearch(tar int, s []int) int {
	return binarySearch(tar, 0, len(s)-1, s)
}

func binarySearch(target int, first int, last int, s []int) int {
	m := math.Round(float64(first + (last-first)/2))
	middle := int(m)
	if first > last {
		panic("no such value in a collection")
	}
	if s[middle] > target {
		return binarySearch(target, first, middle-1, s)
	}
	if s[middle] < target {
		return binarySearch(target, middle+1, last, s)
	}
	return middle
}
