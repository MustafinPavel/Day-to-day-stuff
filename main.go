package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//ввод
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	nailsAmount := readSingleInt(in)
	nails := readShortIntSlice(in)
	sort.Sort(sort.IntSlice(nails))
	//логика
	memo := make([]int, nailsAmount)
	for i := 0; i < nailsAmount; i++ {
		memo[i] = -1
	}
	answer := dp(nails, memo, nailsAmount-1)
	//вывод
	out.WriteString(strconv.Itoa(answer))
}
func dp(nails, memo []int, i int) int {
	if i == 0 || i == 1 {
		return nails[1] - nails[0]
	}
	if memo[i] != -1 {
		return memo[i]
	}
	option1 := dp(nails, memo, i-2) + nails[i] - nails[i-1]
	option2 := dp(nails, memo, i-1) + nails[i] - nails[i-1]
	if option1 <= option2 {
		memo[i] = option1
		return option1
	}
	memo[i] = option2
	return option2

}

func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
func readShortIntSlice(r *bufio.Reader) []int {
	result := make([]int, 0, 0)
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}
