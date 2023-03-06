package main

import (
	"bufio"
	"math"
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
	timeA := readTime(in)
	timeB := readTime(in)
	timeC := readTime(in)
	//логика
	var fullDay int = 24 * 60 * 60
	var difference int
	if timeA > timeC {
		timeC += fullDay
	}
	difference = timeC - timeA

	correctServerTime := (difference / 2) + timeB
	if math.Remainder(float64(difference), 2) >= 0.5 {
		correctServerTime += 1
	}
	if correctServerTime > fullDay {
		correctServerTime -= fullDay
	}
	//вывод
	result := make([]int, 3, 3)
	result[0] = correctServerTime / 3600
	result[1] = (correctServerTime - result[0]*3600) / 60
	result[2] = correctServerTime % 60
	for i, time := range result {
		if time < 10 {
			out.WriteString("0")
		}
		out.WriteString(strconv.Itoa(time))
		if i != 2 {
			out.WriteString(":")
		}
	}
}
func readTime(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	str := strings.Split(string(line), ":")
	hours, _ := strconv.Atoi(str[0])
	minutes, _ := strconv.Atoi(str[1])
	secs, _ := strconv.Atoi(str[2])
	return secs + minutes*60 + hours*3600
}
