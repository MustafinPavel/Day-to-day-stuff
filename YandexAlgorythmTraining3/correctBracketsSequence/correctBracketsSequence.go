package correctBracketsSequence

import (
	"bufio"
	"os"
)

// Checks if the brackets sequence is correct ()[]{}
func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	bracketLine := readLongLine(in)
	//logic
	var stack Stack
	opened := 0
	closed := 0
	itsFine := true
	for i := 0; i < len(bracketLine); i++ {
		this := string(bracketLine[i])
		switch {
		case this == "(" || this == "{" || this == "[":
			stack.Push(this)
			opened++
		case this == ")":
			pop := stack.Pop()
			closed++
			if pop != "(" {
				itsFine = false
			}
		case this == "]":
			pop := stack.Pop()
			closed++
			if pop != "[" {
				itsFine = false
			}
		case this == "}":
			pop := stack.Pop()
			closed++
			if pop != "{" {
				itsFine = false
			}
		}
		if !itsFine {
			break
		}
	}
	if itsFine && opened == closed {
		out.WriteString("yes")
	} else {
		out.WriteString("no")
	}
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

type Stack struct {
	stack []string
	len   int
}

func (s *Stack) Push(n string) {
	s.stack = append(s.stack, n)
	s.len++
}
func (s *Stack) Pop() (n string) {
	if s.len != 0 {
		s.len--
		n = s.stack[len(s.stack)-1]
		s.stack = s.stack[:s.len]
	}
	return
}
