package usefulFuncs

type Stack struct {
	stack []int
	len   int
}

func (s *Stack) Push(n int) {
	s.stack = append(s.stack, n)
	s.len++
}

func (s *Stack) Pop() int {
	if s.len != 0 {
		result := s.stack[len(s.stack)-1]
		s.len--
		s.stack = s.stack[:s.len]
		return result
	} else {
		panic("Stack is empty")
	}
}

func (s *Stack) Back() int {
	if s.len != 0 {
		return s.stack[len(s.stack)-1]
	} else {
		panic("Stack is empty")
	}
}

func (s *Stack) Size() int {
	return s.len
}

func (s *Stack) Clear() {
	s.stack = make([]int, 0)
	s.len = 0
}
