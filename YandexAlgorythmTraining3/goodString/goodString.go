package goodString

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	//ввод
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	symbAmount := readSingleInt(in)
	dictionary := make([]int, 0, symbAmount)
	for i := 0; i < symbAmount; i++ {
		newSymb := readSingleInt(in)
		dictionary = append(dictionary, newSymb)
	}
	// логика
	var result int
	for i := 0; i < symbAmount-1; i++ {
		result += chooseMin(dictionary[i], dictionary[i+1])
	}
	//вывод
	out.WriteString(strconv.Itoa(result))

}

func chooseMin(a, b int) int {
	if a >= b {
		return b
	}
	return a
}

func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
