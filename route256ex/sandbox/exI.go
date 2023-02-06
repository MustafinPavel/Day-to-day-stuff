package main

// import (
// 	"bufio"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	file1, _ := os.Open("testFileInput")
// 	r := bufio.NewReader(file1)
// 	// r := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()
// 	procMissionsCount := ScanLine(r)
// 	//procSlice := ScanLine(r)
// 	var missionsArray [][]int = make([][]int, 0, procMissionsCount[1])
// 	for i := 0; i < procMissionsCount[1]; i++ {
// 		missionsArray[i] = ScanLine(r)
// 	}
// 	for i := 0; i < procMissionsCount[1]; i++ {

// 	}

// }
// func ScanLine(r *bufio.Reader) []int {
// 	line, _, _ := r.ReadLine()
// 	stringsInLine := strings.Fields(string(line))
// 	intSlice := make([]int, 0, len(stringsInLine)*2)
// 	for _, v := range stringsInLine {
// 		num, _ := strconv.Atoi(v)
// 		intSlice = append(intSlice, num)
// 	}
// 	return intSlice
// }
