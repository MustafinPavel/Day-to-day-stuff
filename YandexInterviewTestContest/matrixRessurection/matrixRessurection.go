//Недостаточно быстрое решение. Баг с мемоизацией.

package matrixRessurection

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	tab := readTable(in)
	//main logic
	memory := make(map[[2]int]int)
	for i := 0; i < len(tab); i++ {
		for j := 0; j < len(tab[0]); j++ {
			checkMaxLength(i, j, tab, 0, memory)
		}
	}
	result := 0
	for _, v := range memory {
		if v > result {
			result = v
		}
	}
	out.WriteString(strconv.Itoa(result))
}

func checkMaxLength(i, j int, tab [][]int, currentDepth int, memory map[[2]int]int) {
	checkOneNeighbour(i-1, j, tab, tab[i][j], currentDepth, memory)
	checkOneNeighbour(i+1, j, tab, tab[i][j], currentDepth, memory)
	checkOneNeighbour(i, j-1, tab, tab[i][j], currentDepth, memory)
	checkOneNeighbour(i, j+1, tab, tab[i][j], currentDepth, memory)
	memory[[2]int{i, j}] = int(math.Max(float64(memory[[2]int{i, j}]), float64(currentDepth+1)))
}

func checkOneNeighbour(i, j int, tab [][]int, currentValue int, currentDepth int, memory map[[2]int]int) {
	if isExist(i, j, tab) && tab[i][j] < currentValue {
		// if checkInMemory(i, j, memory) {
		// 	currentDepth += memory[[2]int{i, j}]
		// } else {
		checkMaxLength(i, j, tab, currentDepth+1, memory)
		// }
	}
}

func checkInMemory(i, j int, memory map[[2]int]int) bool {
	if _, ok := memory[[2]int{i, j}]; ok {
		return true
	}
	return false
}
func isExist(i, j int, tab [][]int) bool {
	if i >= 0 && i < len(tab) && j >= 0 && j < len(tab[0]) {
		return true
	}
	return false
}

// Data input
func readTable(r *bufio.Reader) (result [][]int) {
	tableSize := readSliceFromLine(r)
	for i := 0; i < tableSize[0]; i++ {
		tableLine := readSliceFromLine(r)
		result = append(result, tableLine)
	}
	return
}
func readSliceFromLine(r *bufio.Reader) (result []int) {
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return
}
