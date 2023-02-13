package status200

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// This function accepts a slice and counts, how many unique pairs |a-b|%200 are there
func Solution() {
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	numsAmount := readOneInt(in)
	numLine := readFullLine(in)
	nums := parseLineToInts(numLine)
	result := 0
	resultingMap := make(map[int]int)

	//Второе оптимизированное решение. Временная сложность чуть больше O(2N)
	for i := 0; i < numsAmount; i++ {
		tmp := nums[i] % 200
		if tmp >= 0 {
			resultingMap[tmp] += 1
		} else {
			resultingMap[-tmp] += 1
		}
	}
	for _, v := range resultingMap {
		tmp := 0
		if v > 1 {
			for i := 1; i < v; i++ {
				tmp += i
			}
		}
		result += tmp
	}
	out.WriteString(strconv.Itoa(result))
}

// Reading and converting input stream:
func readOneInt(in *bufio.Reader) int {
	temp, _, _ := in.ReadLine()
	tempInt, _ := strconv.Atoi(string(temp))
	return tempInt
}
func readFullLine(in *bufio.Reader) []byte {
	isNotEnded := true
	byteSlice := make([]byte, 0, 100000)
	for isNotEnded {
		cds, end := readLinePart(in)
		isNotEnded = end
		byteSlice = append(byteSlice, cds...)
	}
	return byteSlice
}
func readLinePart(in *bufio.Reader) ([]byte, bool) {
	temp, isNotEnded, _ := in.ReadLine()
	return temp, isNotEnded
}
func parseLineToInts(in []byte) []int {
	tmpStrSlice := strings.Split(string(in), " ")
	nums := make([]int, 0, len(tmpStrSlice)*2)
	for _, v := range tmpStrSlice {
		num, _ := strconv.Atoi(v)
		nums = append(nums, num)
	}
	return nums
}
