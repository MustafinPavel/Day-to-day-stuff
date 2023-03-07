package sumInRectangle

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Решение с полной мемоизацией всех значений сумм прямоугольников
func main() {
	//ввод
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	xSize, _, requestsAmount := readMatrixSize(in)
	matrix := make([][]int, 0, xSize)
	memory := make([][]int, 0, xSize)
	requests := make([][]int, 0, requestsAmount)
	for i := 0; i < xSize; i++ {
		row := readMatrixRow(in)
		matrix = append(matrix, row)
	}
	for i := 0; i < requestsAmount; i++ {
		row := readMatrixRow(in)
		requests = append(requests, row)
	}
	for x := 0; x < len(matrix); x++ {
		accumulatedSums := make([]int, 0, len(matrix[0]))
		tmp := 0
		if x == 0 {
			for y := 0; y < len(matrix[0]); y++ {
				tmp += matrix[x][y]
				accumulatedSums = append(accumulatedSums, tmp)
			}
		} else {
			for y := 0; y < len(matrix[0]); y++ {
				if y == 0 {
					tmp = tmp + memory[x-1][y] + matrix[x][y]
				} else {
					tmp = memory[x-1][y] + accumulatedSums[y-1] - memory[x-1][y-1] + matrix[x][y]
				}
				accumulatedSums = append(accumulatedSums, tmp)
			}
		}
		memory = append(memory, accumulatedSums)
		tmp = 0
	}
	// логика
	var sum int
	for _, req := range requests {
		switch {
		case req[0] == 1 && req[1] == 1:
			sum = memory[req[2]-1][req[3]-1]
		case req[0] == 1:
			sum = memory[req[2]-1][req[3]-1] - memory[req[2]-1][req[1]-2]
		case req[1] == 1:
			sum = memory[req[2]-1][req[3]-1] - memory[req[0]-2][req[3]-1]
		default:
			sum = memory[req[2]-1][req[3]-1] - memory[req[2]-1][req[1]-2] - memory[req[0]-2][req[3]-1] + memory[req[0]-2][req[1]-2]
		}
		//вывод
		out.WriteString(strconv.Itoa(sum) + "\n")
		sum = 0
	}
}

func readMatrixSize(r *bufio.Reader) (x, y, requests int) {
	line, _, _ := r.ReadLine()
	str := strings.Split(string(line), " ")
	x, _ = strconv.Atoi(str[0])
	y, _ = strconv.Atoi(str[1])
	requests, _ = strconv.Atoi(str[2])
	return
}
func readMatrixRow(in *bufio.Reader) []int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 10)
	result := make([]int, 0, 10)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	slice := strings.Fields(string(tmpByteSlice))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}
