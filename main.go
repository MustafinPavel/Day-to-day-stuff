package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	countriesAmount := readSingleInt(in)
	countries := readLongIntSlice(in)
	fmt.Println(countries, countriesAmount) //plug

}

type Stack struct {
	stack []int
	len   int
}

func (s *Stack) Push(n int) {
	s.stack = append(s.stack, n)
	s.len++
}

func (s *Stack) Pop() (n int) {
	if s.len != 0 {
		s.len--
		n = s.stack[len(s.stack)-1]
		s.stack = s.stack[:s.len]
	}
	return
}
func (s *Stack) ShowLast() (n int) {
	return s.stack[len(s.stack)-1]
}

func readLongIntSlice(in *bufio.Reader) []int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	slice := strings.Fields(string(tmpByteSlice))
	result := make([]int, 0, len(slice))
	for _, v := range slice {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}
func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
