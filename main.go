package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//ввод
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	str := readShortLine(in)
	fmt.Printf("%v, %v", str, len(str))
}
func readShortLine(r *bufio.Reader) string {
	line, _, _ := r.ReadLine()
	return string(line)
}
