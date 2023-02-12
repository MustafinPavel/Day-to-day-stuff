package exA

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Solution() {
	// //построчный input
	// file1, _ := os.Open("input.txt")
	// in := bufio.NewScanner(file1)
	// in.Scan()
	// jewels := in.Text()
	// //output
	// out := bufio.NewWriter(os.Stdout)
	// defer out.Flush()

	//file1, _ := os.Open("testFileInput")
	//r := bufio.NewReader(file1)
	r := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	linesAmount := readIntFromLine(r)
	for i := 0; i < linesAmount; i++ {
		t := readSliceFromLine(r)
		fmt.Println(t[0] + t[1])
	}
}

func readIntFromLine(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}

func readSliceFromLine(r *bufio.Reader) []int {
	result := make([]int, 0, 100000)
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}
