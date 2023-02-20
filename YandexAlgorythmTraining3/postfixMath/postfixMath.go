package postfixMath

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Counts the result of a polish/postfix-written math operations
func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var stack Stack
	fullQuery := readLongSlice(in)
	for _, v := range fullQuery {
		num, err := strconv.Atoi(v)
		if err != nil {
			last := stack.Pop()
			prelast := stack.Pop()
			switch v {
			case "-":
				stack.Push(prelast - last)
			case "+":
				stack.Push(prelast + last)
			case "*":
				stack.Push(prelast * last)
			}
		} else {
			stack.Push(num)
		}
	}
	out.WriteString(strconv.Itoa(stack.Pop()))

}
func readLongSlice(in *bufio.Reader) []string {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	slice := strings.Fields(string(tmpByteSlice))
	return slice
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
