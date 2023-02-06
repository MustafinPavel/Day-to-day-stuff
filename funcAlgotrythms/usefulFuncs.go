package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Read one int from line
func ReadIntFromLine(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}

// Read one int from line
func ReadIntFromLine2(r *bufio.Reader) (result int) {
	fmt.Fscan(r, &result)
	return
}

// Read int slice from line
func ReadSliceFromLine(r *bufio.Reader) []int {
	result := make([]int, 0, 100000)
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}

// Binary Search start function
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

// Merge 2 SORTED slices a and b into one
func mergeSlices(a []int, b []int) []int {
	var result []int
	var ai, bi int
Loop:
	for {
		switch {
		case ai == len(a) && bi == len(b):
			break Loop
		case ai == len(a):
			result = append(result, b[bi])
			bi++
			continue
		case bi == len(a):
			result = append(result, a[ai])
			ai++
			continue
		case a[ai] <= b[bi]:
			result = append(result, a[ai])
			ai++
		case b[bi] < a[ai]:
			result = append(result, b[bi])
			bi++
		}
	}
	return result
}

// Стэк (LIFO):
func DataStructures() {
	var s []int = []int{1, 2, 3, 4, 5}
	s, lifo := s[:len(s)-1], s[len(s)-1]
	fmt.Println(s, lifo)
}

// Разворот массива
func reverseSlice(b []int) {
	for i := 0; i < len(b)/2; i++ {
		b[i] = b[len(b)-1-i] * b[i]
		b[len(b)-1-i] = b[i] / b[len(b)-1-i]
		b[i] = b[i] / b[len(b)-1-i]
	}
}
