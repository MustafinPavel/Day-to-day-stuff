package exD

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Неплохая структура
func Solution() {
	// file1, _ := os.Open("testFileInput")
	// r := bufio.NewReader(file1)
	r := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	tablesAmount := readIntFromLine(r)
	for i := 0; i < tablesAmount; i++ {
		table := readTable(r)
		clicks := readClicks(r)
		for _, v := range clicks[1] {
			table = mySort(table, v-1)
		}
		printFinalTable(table)
		fmt.Println()
	}
}
func mySort(table [][]int, clickOn int) [][]int {
	sort.SliceStable(table, func(i, j int) bool {
		return table[i][clickOn] < table[j][clickOn]
	})
	return table
}
func printFinalTable(table [][]int) {
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
func readIntFromLine(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
func readSliceFromLine(r *bufio.Reader) (result []int) {
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return
}
func readTable(r *bufio.Reader) (result [][]int) {
	r.ReadLine() //skip empty line
	keyLine := readSliceFromLine(r)
	for i := 0; i < keyLine[0]; i++ {
		tableLine := readSliceFromLine(r)
		result = append(result, tableLine)
	}
	return
}
func readClicks(r *bufio.Reader) (result [][]int) {
	for i := 0; i < 2; i++ {
		clicksLine := readSliceFromLine(r)
		result = append(result, clicksLine)
	}
	return
}
