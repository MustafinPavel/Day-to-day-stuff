package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
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

//____________________

// type Stack struct {
// 	stack []string
// }
// func (st *Stack) AddToStack(s string) {
// 	st.stack = append(st.stack, s)
// }
// func (st *Stack) PopFromStack() string {
// 	if len(st.stack) != 0 {
// 		returningString := st.stack[len(st.stack)-1]
// 		st.stack = st.stack[:len(st.stack)-1]
// 		return returningString
// 	}
// 	panic("The stack is empty")
// }

// func InsertionSort(array []int) {
// 	for i := 0; i < len(array); i++ {
// 		j := i
// 		temp := array[i]
// 		for j > 0 && array[j-1] > temp {
// 			array[j] = array[j-1]
// 			j--
// 		}
// 		array[j] = temp
// 	}
// }
