package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func exD() {
	// file1, _ := os.Open("testFileInput")
	// r := bufio.NewReader(file1)
	r := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	tablesAmount := ReadIntFromLine(r)
	for i := 0; i < tablesAmount; i++ {
		table := ReadTable(r)
		clicks := ReadClicks(r)
		for _, v := range clicks[1] {
			table = Sort(table, v-1)
		}
		PrintFinalTable(table)
		fmt.Println()
	}
}
func Sort(table [][]int, clickOn int) [][]int {
	sort.SliceStable(table, func(i, j int) bool {
		return table[i][clickOn] < table[j][clickOn]
	})
	return table
}
func PrintFinalTable(table [][]int) {
	for i := 0; i < len(table); i++ {
		for k, v := range table[i] {
			if k == (len(table[i]) - 1) {
				fmt.Printf("%d\n", v)
			} else {
				fmt.Printf("%d ", v)
			}
		}
	}
}
func ReadIntFromLine(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
func ReadSliceFromLine(r *bufio.Reader) (result []int) {
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return
}
func ReadTable(r *bufio.Reader) (result [][]int) {
	r.ReadLine() //skip empty line
	keyLine := ReadSliceFromLine(r)
	for i := 0; i < keyLine[0]; i++ {
		tableLine := ReadSliceFromLine(r)
		result = append(result, tableLine)
	}
	return
}
func ReadClicks(r *bufio.Reader) (result [][]int) {
	for i := 0; i < 2; i++ {
		clicksLine := ReadSliceFromLine(r)
		result = append(result, clicksLine)
	}
	return
}
