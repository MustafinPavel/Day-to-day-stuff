package main

import (
	"bufio"
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
	fieldSize := readShortIntSlice(in)
	rows, columns := fieldSize[0]+1, fieldSize[1]+1
	field := make([][]int, rows+1)
	emptyLine := make([]int, columns)
	for i := 0; i < len(emptyLine); i++ {
		emptyLine[i] = 999
	}
	field[0] = emptyLine
	for i := 1; i <= rows; i++ {
		line := readShortIntSlice(in)
		field[i] = append(field[i], 999)
		for j := 0; j < len(line); j++ {
			field[i] = append(field[i], line[j])
		}
	}
	//логика
	answer := dp(field, columns-1, rows-1)
	//вывод
	out.WriteString(strconv.Itoa(answer))
}

func dp(field [][]int, posX, posY int) int {
	switch {
	case posX == 1 && posY == 1:
		return field[posY][posX]
	case posX == 1:
		option := dp(field, posX, posY-1) + field[posY][posX]
		return option
	case posY == 1:
		option := dp(field, posX-1, posY) + field[posY][posX]
		return option
	default:
		option2 := dp(field, posX, posY-1) + field[posY][posX]
		option1 := dp(field, posX-1, posY) + field[posY][posX]
		if option1 < option2 {
			return option1
		}
		return option2
	}
}

func readShortIntSlice(r *bufio.Reader) []int {
	result := make([]int, 0, 0)
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}
