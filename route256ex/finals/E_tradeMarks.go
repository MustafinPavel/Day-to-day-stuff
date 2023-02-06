// ПРИНЯТА НЕ ПОЛНОСТЬЮ
package main

// import (
// 	"bufio"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// //Выбирает из списка недублирующиеся товарные знаки с учётом малых отличий
// func main() {
// 	r := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()
// 	setsAmount := ReadIntFromLine(r)
// 	for i := 0; i < setsAmount; i++ {
// 		namesAmount := ReadIntFromLine(r)
// 		allNamesWithCoresSlice := make([][]string, 0, 1000000)
// 		for j := 0; j < namesAmount; j++ {
// 			scannedName := ReadOrdinaryString(r)
// 			scannedNameCore := GetCore(scannedName)
// 			allNamesWithCoresSlice = append(allNamesWithCoresSlice, []string{scannedName, scannedNameCore})
// 		}
// 		numberOfRegisteredTitles := CompareCoresAndSelectResults(allNamesWithCoresSlice)
// 		out.WriteString(strconv.Itoa(numberOfRegisteredTitles) + "\n")
// 	}
// }
// func CompareCoresAndSelectResults(sl [][]string) (result int) {
// 	checkedItems := make(map[int]bool)
// 	for i := 0; i < len(sl); i++ {
// 		if _, ok := checkedItems[i]; !ok {
// 			result++
// 		}
// 		for j := 0; j < len(sl); j++ {
// 			if sl[i][1] == sl[j][1] && i != j {
// 				if areSimilar(sl[i][0], sl[j][0], len(sl[i][1])) {
// 					checkedItems[j] = true
// 				}
// 			}
// 		}
// 	}
// 	return
// }
// func areSimilar(strA string, strB string, coreLength int) bool {
// 	var tallyA, tallyB int
// 	compatisonSliceA := make([]int, 0, 10000)
// 	compatisonSliceB := make([]int, 0, 10000)

// 	for i := 0; i < len(strA); i++ {
// 		if i == 0 || strA[i] == strA[i-1] {
// 			tallyA++
// 			continue
// 		}
// 		if strA[i] != strA[i-1] || i == len(strA)-1 {
// 			compatisonSliceA = append(compatisonSliceA, tallyA)
// 			tallyA = 1
// 		}
// 	}
// 	compatisonSliceA = append(compatisonSliceA, tallyA)
// 	for i := 0; i < len(strB); i++ {
// 		if i == 0 || strB[i] == strB[i-1] {
// 			tallyB++
// 			continue
// 		}
// 		if strB[i] != strB[i-1] || i == len(strB)-1 {
// 			compatisonSliceB = append(compatisonSliceB, tallyB)
// 			tallyB = 1
// 		}
// 	}
// 	compatisonSliceB = append(compatisonSliceB, tallyB)
// 	truenessCounter := 0
// 	for k, v := range compatisonSliceA {
// 		if v == compatisonSliceB[k] || v > 1 && compatisonSliceB[k] > 1 {
// 			truenessCounter++
// 		}
// 	}
// 	return truenessCounter == coreLength
// }
// func GetCore(input string) string {
// 	result := ""
// 	for k, v := range input {
// 		if k == 0 {
// 			result += string(v)
// 			continue
// 		}
// 		if input[k] == input[k-1] {
// 			continue
// 		} else {
// 			result += string(v)
// 		}
// 	}
// 	return result
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

// func ReadOrdinaryString(r *bufio.Reader) string {
// 	line, _, _ := r.ReadLine()
// 	return string(line)
// }
