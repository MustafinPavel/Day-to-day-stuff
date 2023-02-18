package stockTrading

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solution is in Another Castle.
func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	numOfDays := readLineWithOneInt(in)
	allData := readIntSlice(in)
	//logic
	fmt.Println(numOfDays, allData) //plug (DELETE)

}

func readIntSlice(in *bufio.Reader) []int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 300000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	line := string(tmpByteSlice)
	result := make([]int, 0, len(line)/2)
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}

func readLineWithOneInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
