package greatLinelandMigration

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Выбирает в исходном массиве ближайшее число справа, которое меньше,
// чем текущее, и помещает его индекс в поле [2]result
func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	countriesAmount := readSingleInt(in)
	countries := readLongIntSlice(in) //[0]index, [1]value, [2]result
	//logic
	var myStack Stack
	for i := 0; i < countriesAmount; i++ {
		next := countries[i]
		for myStack.len > 0 {
			prev := myStack.ShowLast()
			if prev[1] <= next[1] {
				myStack.Push(next)
				break
			} else {
				p := myStack.Pop()
				countries[p[0]][2] = next[0]
			}
		}
		if myStack.len == 0 {
			myStack.Push(next)
		}
	}
	for myStack.len > 0 {
		p := myStack.Pop()
		countries[p[0]][2] = -1
	}
	//output
	for i := 0; i < len(countries); i++ {
		add := countries[i][2]
		out.WriteString(strconv.Itoa(add))
		if i < len(countries)-1 {
			out.WriteString(" ")
		}
	}
}

type Stack struct {
	stack [][]int
	len   int
}

func (s *Stack) Push(n []int) {
	s.stack = append(s.stack, n)
	s.len++
}

func (s *Stack) Pop() (n []int) {
	if s.len != 0 {
		s.len--
		n = s.stack[len(s.stack)-1]
		s.stack = s.stack[:s.len]
	}
	return
}
func (s *Stack) ShowLast() (n []int) {
	return s.stack[len(s.stack)-1]
}

func readLongIntSlice(in *bufio.Reader) [][]int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	slice := strings.Fields(string(tmpByteSlice))
	result := make([][]int, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		j, _ := strconv.Atoi(slice[i])
		result = append(result, []int{i, j, 0})
	}
	return result
}
func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
