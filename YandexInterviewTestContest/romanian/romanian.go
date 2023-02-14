package romanian

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Проверка корректности введённого римского числа и выдача его арабской записью
func main() {
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	romanString := readOrdinaryString(in)
	romanInts := make([]int, len(romanString))
	for i := 0; i < len(romanString); i++ {
		romanInts[i] = romanLetterToInt(string(romanString[i]))
	}
	ok := checkForInputErrors(romanString, romanInts)
	if !ok || len(romanString) == 0 {
		out.WriteString("-1")
	} else {
		resultNumsRow := make([]int, 0, len(romanInts))
		result := 0
		for i := 0; i < len(romanInts); i++ {
			current := i
			next := i + 1
			if next < len(romanInts) && romanInts[next] <= romanInts[current] {
				plus := romanInts[current]
				result += plus
				resultNumsRow = append(resultNumsRow, plus)
			}
			if next < len(romanInts) && romanInts[next] > romanInts[current] {
				plus := romanInts[next] - romanInts[current]
				result += plus
				resultNumsRow = append(resultNumsRow, plus)
				i++
			}
			if next == len(romanInts) {
				plus := romanInts[current]
				result += plus
				resultNumsRow = append(resultNumsRow, plus)
			}
		}
		if isDescending(resultNumsRow) { //checking the descending pattern of a number
			out.WriteString(strconv.Itoa(result))
		} else {
			out.WriteString("-1")
		}
	}
}

func readOrdinaryString(r *bufio.Reader) string {
	line, _, _ := r.ReadLine()
	return string(line)
}

func checkForInputErrors(s string, ints []int) bool {
	if strings.Count(s, "V") > 1 || strings.Count(s, "L") > 1 || strings.Count(s, "D") > 1 {
		return false
	}
	seqLetCounter := 1
	for i := 0; i < len(s); i++ {
		if i > 0 {
			if s[i] == s[i-1] {
				t := string(s[i])
				if t == "V" || t == "L" || t == "D" || seqLetCounter == 3 {
					return false
				} else {
					seqLetCounter++
				}
			} else {
				seqLetCounter = 1
			}
			if ints[i] > ints[i-1] {
				if ints[i-1] == 1 {
					if ints[i] != 5 && ints[i] != 10 {
						return false
					}
				}
				if ints[i-1] == 10 {
					if ints[i] != 50 && ints[i] != 100 {
						return false
					}
				}
				if ints[i-1] == 100 {
					if ints[i] != 500 && ints[i] != 1000 {
						return false
					}
				}
				if ints[i-1] != 1 && ints[i-1] != 10 && ints[i-1] != 100 {
					return false
				}
			}
		}
	}
	return true
}

func isDescending(row []int) bool {
	for i := 0; i < len(row); i++ {
		cur := i
		prev := i - 1
		if i > 0 && row[prev] < row[cur] {
			return false
		}
	}
	return true
}

func romanLetterToInt(s string) (res int) {
	switch s {
	case "I":
		res = 1
	case "V":
		res = 5
	case "X":
		res = 10
	case "L":
		res = 50
	case "C":
		res = 100
	case "D":
		res = 500
	case "M":
		res = 1000
	}
	return
}
