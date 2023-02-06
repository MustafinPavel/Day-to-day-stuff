// НЕ ДОДЕЛАНА
package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	"strings"
// 	"unicode"
// )

// func main() {
// 	//ALPH: 65-90
// 	//DIG	48-57
// 	file1, _ := os.Open("input.txt")
// 	r := bufio.NewReader(file1)
// 	//r := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()
// 	numOfStrings := ReadIntFromLine(r)
// 	for i := 0; i < numOfStrings; i++ {
// 		line, _, _ := r.ReadLine()
// 		if _, ok := CheckIfCorrect(line); ok {
// 			fmt.Println(line)
// 		}

// 	}
// }
// func CheckIfCorrect(line []byte) (string, bool) {
// 	result := ""
// 	for i,v:=range line{
// 		if unicode.IsLetter()
// 	}

// 	return result, true

// }
// func ReadIntFromLine(r *bufio.Reader) int {
// 	line, _, _ := r.ReadLine()
// 	lineInt, _ := strconv.Atoi(string(line))
// 	return lineInt
// }
// func ReadSliceFromLine(r *bufio.Reader) []int {
// 	result := make([]int, 0, 100000)
// 	line, _, _ := r.ReadLine()
// 	slice := strings.Fields(string(line))
// 	for i := 0; i < len(slice); i++ {
// 		t, _ := strconv.Atoi(slice[i])
// 		result = append(result, t)
// 	}
// 	return result
// }
