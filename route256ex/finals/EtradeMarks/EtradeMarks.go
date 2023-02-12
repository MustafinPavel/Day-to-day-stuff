// Принята не полностью
package EtradeMarks

import (
	"bufio"
	"os"
	"strconv"
)

// Выбирает из списка недублирующиеся товарные знаки с учётом малых отличий
func SelectUniqueTrademarks() {
	r := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	setsAmount := readIntFromLine(r)
	for i := 0; i < setsAmount; i++ {
		namesAmount := readIntFromLine(r)
		allNamesWithCoresSlice := make([][]string, 0, 1000000)
		for j := 0; j < namesAmount; j++ {
			scannedName := readOrdinaryString(r)
			scannedNameCore := getCore(scannedName)
			allNamesWithCoresSlice = append(allNamesWithCoresSlice, []string{scannedName, scannedNameCore})
		}
		numberOfRegisteredTitles := compareCoresAndSelectResults(allNamesWithCoresSlice)
		out.WriteString(strconv.Itoa(numberOfRegisteredTitles) + "\n")
	}
}
func compareCoresAndSelectResults(sl [][]string) (result int) {
	checkedItems := make(map[int]bool)
	for i := 0; i < len(sl); i++ {
		if _, ok := checkedItems[i]; !ok {
			result++
		}
		for j := 0; j < len(sl); j++ {
			if sl[i][1] == sl[j][1] && i != j {
				if areSimilar(sl[i][0], sl[j][0], len(sl[i][1])) {
					checkedItems[j] = true
				}
			}
		}
	}
	return
}
func areSimilar(strA string, strB string, coreLength int) bool {
	var tallyA, tallyB int
	compatisonSliceA := make([]int, 0, 10000)
	compatisonSliceB := make([]int, 0, 10000)

	for i := 0; i < len(strA); i++ {
		if i == 0 || strA[i] == strA[i-1] {
			tallyA++
			continue
		}
		if strA[i] != strA[i-1] || i == len(strA)-1 {
			compatisonSliceA = append(compatisonSliceA, tallyA)
			tallyA = 1
		}
	}
	compatisonSliceA = append(compatisonSliceA, tallyA)
	for i := 0; i < len(strB); i++ {
		if i == 0 || strB[i] == strB[i-1] {
			tallyB++
			continue
		}
		if strB[i] != strB[i-1] || i == len(strB)-1 {
			compatisonSliceB = append(compatisonSliceB, tallyB)
			tallyB = 1
		}
	}
	compatisonSliceB = append(compatisonSliceB, tallyB)
	truenessCounter := 0
	for k, v := range compatisonSliceA {
		if v == compatisonSliceB[k] || v > 1 && compatisonSliceB[k] > 1 {
			truenessCounter++
		}
	}
	return truenessCounter == coreLength
}
func getCore(input string) string {
	result := ""
	for k, v := range input {
		if k == 0 {
			result += string(v)
			continue
		}
		if input[k] == input[k-1] {
			continue
		} else {
			result += string(v)
		}
	}
	return result
}
func readIntFromLine(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}

func readOrdinaryString(r *bufio.Reader) string {
	line, _, _ := r.ReadLine()
	return string(line)
}
