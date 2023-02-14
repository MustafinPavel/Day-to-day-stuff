package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	//решение:
	goodline := readLongLine(in)
	mask := make([]int, len(goodline))
	for i, j := 1, 0; i < len(goodline); i++ {
		if i > 0 {
			if areSimilar(j, i, goodline, mask) {
				mask[i], mask[j] = 1, 1
				j = findLeftUndeleted(j, mask)
			} else {
				j = i
			}
		}
	}
	//вывод результата:
	for i := 0; i < len(goodline); i++ {
		if mask[i] == 0 {
			out.WriteString(string(goodline[i]))
		}
	}
}
func areSimilar(prev, cur int, line string, mask []int) bool {
	if strings.EqualFold(string(line[cur]), string(line[prev])) && line[cur] != line[prev] && mask[prev] != 1 {
		return true
	}
	return false
}
func findLeftUndeleted(j int, mask []int) int {
	for mask[j] != 0 {
		j--
		if j <= 0 {
			j = 0
			break
		}
	}
	return j
}
func readLongLine(in *bufio.Reader) string {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	return string(tmpByteSlice)
}
