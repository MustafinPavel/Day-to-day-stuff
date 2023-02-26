package threeOnesInARow

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	//ввод
	timer := time.Now()
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var N int64 = readSingleInt(in)
	//логика
	result := countVariations(N)
	//вывод
	out.WriteString(strconv.Itoa(result))
	fmt.Println(time.Since(timer))
}

func countVariations(n int64) int {
	//база ДП
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 2
	}
	if n == 2 {
		return 4
	}
	//рекуррентность ДП
	result := 0
	variants := math.Pow(2, float64(n))
	for i := variants - 1; i >= variants/2; i-- {
		numInBin := strconv.FormatInt(int64(i), 2)
		if !strings.Contains(numInBin, "111") {
			result++
		}
	}
	result += countVariations(n - 1)
	return result
}

func readSingleInt(r *bufio.Reader) int64 {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return int64(lineInt)
}
