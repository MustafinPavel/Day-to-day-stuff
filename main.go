package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

}
func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
func readShortLine(r *bufio.Reader) string {
	line, _, _ := r.ReadLine()
	return string(line)
}
