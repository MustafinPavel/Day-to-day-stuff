package threeOnesInARow

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
	var N int64 = readSingleInt(in)
	//логика
	var memes []int = make([]int, 40)
	result := countVariations(N, memes)
	//вывод
	out.WriteString(strconv.Itoa(result))
}

func countVariations(n int64, mem []int) int {
	//база ДП
	if mem[n] != 0 {
		return mem[n]
	}
	switch n {
	case 0:
		return 0
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 7
	}
	//рекуррентность ДП
	result := 0
	result = countVariations(n-1, mem) + countVariations(n-2, mem) + countVariations(n-3, mem)
	mem[n] = result
	return result
}

func readSingleInt(r *bufio.Reader) int64 {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return int64(lineInt)
}
