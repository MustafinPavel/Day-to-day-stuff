package trainCarsSorting

import (
	"bufio"
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
	carsAmount := readSingleInt(in)
	wayOne := readLongIntSlice(in)
	wayTwo := make([]int, 0, len(wayOne))
	//logic
	var st Stack
	for i := 0; i < carsAmount; i++ {
		next := wayOne[i]
		for st.len > 0 {
			prev := st.ShowLast()
			if prev < next {
				wayTwo = append(wayTwo, st.Pop())
			} else {
				st.Push(next)
				break
			}
		}
		if st.len == 0 {
			st.Push(next)
		}
	}
	for st.len > 0 {
		wayTwo = append(wayTwo, st.Pop())
	}
	//output
	for i := 0; i < len(wayTwo); i++ {
		if i > 0 && wayTwo[i-1] > wayTwo[i] {
			out.WriteString("NO")
			break
		}
		if len(wayTwo) == 1 || (i == len(wayTwo)-1 && wayTwo[i-1] < wayTwo[i]) {
			out.WriteString("YES")
		}
	}
}

type Stack struct {
	stack []int
	len   int
}

func (s *Stack) Push(n int) {
	s.stack = append(s.stack, n)
	s.len++
}

func (s *Stack) Pop() (n int) {
	if s.len != 0 {
		s.len--
		n = s.stack[len(s.stack)-1]
		s.stack = s.stack[:s.len]
	}
	return
}
func (s *Stack) ShowLast() (n int) {
	return s.stack[len(s.stack)-1]
}

func readLongIntSlice(in *bufio.Reader) []int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	slice := strings.Fields(string(tmpByteSlice))
	result := make([]int, 0, len(slice))
	for _, v := range slice {
		i, _ := strconv.Atoi(v)
		result = append(result, i)
	}
	return result
}
func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
