package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Read one line and returns a string
func ReadOrdinaryString(r *bufio.Reader) string {
	line, _, _ := r.ReadLine()
	return string(line)
}

// Read one int from line
func ReadIntFromLine(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
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
		case bi == len(a):
			result = append(result, a[ai])
			ai++
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

type Stack struct {
	stack []string
}

func (st *Stack) AddToStack(s string) {
	st.stack = append(st.stack, s)
}
func (st *Stack) PopFromStack() string {
	if len(st.stack) != 0 {
		returningString := st.stack[len(st.stack)-1]
		st.stack = st.stack[:len(st.stack)-1]
		return returningString
	}
	panic("The stack is empty")
}

// Разворот массива
func reverseSlice(b []int) {
	for i := 0; i < len(b)/2; i++ {
		b[i] = b[len(b)-1-i] * b[i]
		b[len(b)-1-i] = b[i] / b[len(b)-1-i]
		b[i] = b[i] / b[len(b)-1-i]
	}
}

// Проверка того, имплементировала ли структура ВСЕ методы интерфейса.
// Если добавить метод в интерфейс, не даст скомпилироваться
type SomeInterface interface {
	Method()
}
type Implementation struct{}

func (*Implementation) Method() { fmt.Println("Hello, World!") }

var _ SomeInterface = (*Implementation)(nil) // ← вот эта строчка
