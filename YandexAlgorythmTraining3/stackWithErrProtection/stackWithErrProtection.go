package stackWithErrProtection

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stack implementation
func main() {
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var myStack Stack
	for {
		text := readShortLine(in)
		handleQuery(text, &myStack)
	}
}

type Stack struct {
	stack []int
	len   int
}

func (s *Stack) Push(n int) {
	s.stack = append(s.stack, n)
	s.len++
	fmt.Println("ok")
}
func (s *Stack) Pop() {
	if s.len != 0 {
		fmt.Printf("%v\n", s.stack[len(s.stack)-1])
		s.len--
		s.stack = s.stack[:s.len]

	} else {
		fmt.Println("error")
	}
}
func (s *Stack) Back() {
	if s.len != 0 {
		fmt.Printf("%v\n", s.stack[len(s.stack)-1])
	} else {
		fmt.Println("error")
	}
}
func (s *Stack) Size() {
	fmt.Printf("%v\n", s.len)
}
func (s *Stack) Clear() {
	s.stack = make([]int, 0)
	s.len = 0
	fmt.Println("ok")
}
func (s *Stack) Exit() {
	fmt.Println("bye")
	os.Exit(0)
}
func handleQuery(query string, st *Stack) {
	switch {
	case strings.HasPrefix(query, "push"):
		query = strings.TrimPrefix(query, "push ")
		n, _ := strconv.Atoi(query)
		st.Push(n)
	case strings.HasPrefix(query, "pop"):
		st.Pop()
	case strings.HasPrefix(query, "back"):
		st.Back()
	case strings.HasPrefix(query, "size"):
		st.Size()
	case strings.HasPrefix(query, "clear"):
		st.Clear()
	case strings.HasPrefix(query, "exit"):
		st.Exit()
	case strings.HasPrefix(query, "EOF"):
		os.Exit(1)
	}
}
func readShortLine(r *bufio.Reader) string {
	line, _, err := r.ReadLine()
	if err != nil {
		return "EOF"
	}
	return string(line)
}
