package main

import (
	"fmt"
)

func main() {

}

func MyReverse(s []string) {
	for i := 0; i < len(s)/2; i++ {
		first := s[i]
		last := s[len(s)-1-i]
		s[len(s)-1-i] = first
		s[i] = last
	}
}
func MySplit(s string) []string {
	temp := ""
	result := make([]string, 0, 10)
	for k, v := range s {
		switch {
		case string(v) != " ":
			temp += string(v)
		case string(v) == " ":
			fmt.Println(k)
			result = append(result, temp)
			temp = ""
		}
	}
	result = append(result, temp)
	return result
}

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

func InsertionSort(array []int) {
	for i := 0; i < len(array); i++ {
		j := i
		temp := array[i]
		for j > 0 && array[j-1] > temp {
			array[j] = array[j-1]
			j--
		}
		array[j] = temp
	}
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
