package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	//ввод
	file1, _ := os.Open("input2.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	K := readSingleInt(in)
	S := readLongLine(in)
	// логика
	maxLen := 0
	symbDict := make(map[byte]bool)
	for i := 0; i < len(S); i++ {
		symbDict[S[i]] = true
	}
	for keySymbol := range symbDict {
		rightIdx := 0
		leftIdx := 0
		Kleft := K
		for rightIdx < len(S) {
			if S[rightIdx] == keySymbol || Kleft > 0 {
				prev := S[rightIdx]
				rightIdx++
				if prev != keySymbol {
					Kleft--
				}
			} else {
				prev := S[leftIdx]
				leftIdx++
				if prev != keySymbol {
					Kleft++
				}
			}
			maxLen = maxOf(maxLen, rightIdx-leftIdx)
		}
	}
	//вывод
	out.WriteString(strconv.Itoa(maxLen))

}
func maxOf(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}

func readLongLine(in *bufio.Reader) string {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 1000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	return string(tmpByteSlice)
}
