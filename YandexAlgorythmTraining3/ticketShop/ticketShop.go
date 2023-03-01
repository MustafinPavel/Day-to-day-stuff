package ticketShop

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	//ввод
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	customers := readSingleInt(in)
	prices := make([][]int, 1, customers)
	for i := 0; i < customers; i++ {
		tmp := readShortIntSlice(in)
		prices = append(prices, tmp)
	}
	//логика
	memes := make([]int, customers+1)
	answer := dp(customers, prices, memes)
	//вывод
	out.WriteString(strconv.Itoa(answer))
}

func dp(n int, pr [][]int, mem []int) int {
	//база
	if mem[n] != 0 {
		return mem[n]
	}
	var min int
	switch n {
	case 1:
		mem[n] = pr[1][0]
		return pr[1][0]
	case 2:
		min = pr[1][0] + pr[2][0]
		a := pr[1][1]
		if a < min {
			min = a
		}
		mem[n] = min
		return min
	case 3:
		min = pr[1][0] + pr[2][0] + pr[3][0]
		a := pr[1][1] + pr[3][0]
		b := pr[1][0] + pr[2][1]
		c := pr[1][2]
		if a < min {
			min = a
		}
		if b < min {
			min = b
		}
		if c < min {
			min = c
		}
		mem[n] = min
		return min
	}
	//dp
	if n > 3 {
		min = dp(n-3, pr, mem) + pr[n-2][2]
		a := dp(n-2, pr, mem) + pr[n-1][1]
		if a < min {
			min = a
		}
		b := dp(n-1, pr, mem) + pr[n][0]
		if b < min {
			min = b
		}
	}
	mem[n] = min
	return min
}

func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
func readShortIntSlice(r *bufio.Reader) []int {
	result := make([]int, 0, 0)
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}
