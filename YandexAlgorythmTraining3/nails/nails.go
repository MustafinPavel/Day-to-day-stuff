package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Не завершена. Решение неверное - подумать над dp.
func main() {
	//ввод
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	nailsAmount := readSingleInt(in)
	nails := readShortIntSlice(in)
	sort.Sort(sort.IntSlice(nails))
	distances := make([]int, 0, len(nails))
	for i := 1; i < nailsAmount; i++ {
		distances = append(distances, nails[i]-nails[i-1])
	}
	//логика
	answer := count(0, distances)
	fmt.Println(distances)
	//вывод
	out.WriteString(strconv.Itoa(answer))
}

func count(n int, dist []int) int {
	mem := make([]int, len(dist))
	var res int
	if n == 0 {
		mem[n] = 1
		n++
	}
	for ; n < len(dist); n++ {
		if n == len(dist)-1 || n == len(dist)-2 {
			mem[len(dist)-1] = 1
			break
		}
		cur := dist[n]
		next := dist[n+1]
		if cur < next {
			mem[n] = 1
		} else {
			mem[n+1] = 1
			n++
		}
	}
	for i, tmp := 0, 0; i < len(mem); i++ {
		if mem[i] == 1 {
			tmp++
		}
		if tmp == 3 {
			tmp = 0
			mem[i-1] = 0
		}
	}
	for i := 0; i < len(mem); i++ {
		if mem[i] == 1 {
			res += dist[i]
		}
	}
	return res
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
