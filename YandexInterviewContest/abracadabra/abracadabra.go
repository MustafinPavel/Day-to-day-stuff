package abracadabra

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Extreeeeeemely slow algorythm, but it works with average data inputs :-/
// Optimization reqired
func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	symbolsInLine, maxSpellLen := readLineWithTwoInts(in)
	symbols := readLongLine(in)
	nextSymbPointers := readIntSlice(in)
	shiftsLine := readIntSlice(in)
	letters := make(map[string]int)
	letters2 := make(map[int]string)
	for i := 0; i < 26; i++ {
		r := rune(97 + i)
		letters[string(r)] = i
		letters2[i] = string(r)
	}
	fmt.Println(symbolsInLine, maxSpellLen, symbols, nextSymbPointers, shiftsLine) //DELETE
	//logic
	UltimatePower := 0
	for i := 0; i < symbolsInLine; i++ {
		symbolsPositions := make(map[int]int)
		totalSymbolsInSpell := make(map[string]int)
		currentPos := i
		for j := 0; j < maxSpellLen; j++ {
			symbolsPositions[currentPos] += 1
			currentSym := string(symbols[currentPos])
			if symbolsPositions[currentPos] > 1 {
				symbolPos := (letters[currentSym] + ((symbolsPositions[currentPos] - 1) * shiftsLine[currentPos])) % 26
				currentSym = letters2[symbolPos]
				totalSymbolsInSpell[currentSym] += 1
			} else {
				totalSymbolsInSpell[currentSym] += 1
			}
			UltimatePower += len(totalSymbolsInSpell)
			currentPos = nextSymbPointers[currentPos] - 1
		}
	}
	fmt.Printf("Ultimate Power: %v\n", UltimatePower)
}

func readLineWithTwoInts(r *bufio.Reader) (n, k int) {
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	n, _ = strconv.Atoi(slice[0])
	k, _ = strconv.Atoi(slice[1])
	return
}
func readIntSlice(in *bufio.Reader) []int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 300000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	line := string(tmpByteSlice)
	result := make([]int, 0, len(line)/2)
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}
func readLongLine(in *bufio.Reader) string {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	return string(tmpByteSlice)
}
