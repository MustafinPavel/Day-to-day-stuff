package betterKeyboard

import (
	"bufio"
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
	numberOfButtons := readLineWithOneInt(in)
	buttonKeys := readIntSlice(in)
	buttonRows := readIntSlice(in)
	numberOfLetters := readLineWithOneInt(in)
	qualificationWorkText := readIntSlice(in)
	//logic
	buttonsMap := make(map[int]int)
	fillTheMap(numberOfButtons, buttonKeys, buttonRows, buttonsMap)
	jumpCounter := 0
	for i := 0; i < numberOfLetters; i++ {
		if i > 0 {
			prev := qualificationWorkText[i-1]
			current := qualificationWorkText[i]
			if buttonsMap[current] != buttonsMap[prev] {
				jumpCounter++
			}
		}
	}
	//output
	out.WriteString(strconv.Itoa(jumpCounter))
}
func fillTheMap(num int, keys []int, rows []int, mp map[int]int) {
	for k, v := range keys {
		mp[v] = rows[k]
	}

}
func readIntSlice(in *bufio.Reader) []int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
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
