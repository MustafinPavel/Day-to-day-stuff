package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	opsNumber := readSingleInt(in)
	goodsInTrain := make(map[string]int)
	var opsNums StackInt
	var opsNames StackStr
	for i := 0; i < opsNumber; i++ {
		line := readShortLine(in)
		if strings.HasPrefix(line, "get") {
			trimmed := strings.TrimPrefix(line, "get ")
			if amount, ok := goodsInTrain[trimmed]; ok {
				out.WriteString(strconv.Itoa(amount) + "\n")
			} else {
				out.WriteString("0" + "\n")
			}
		}
		if strings.HasPrefix(line, "add") {
			tmp := strings.Fields(line)
			amount, _ := strconv.Atoi(tmp[1])
			name := tmp[2]
			opsNames.Push(name)
			opsNums.Push(amount)
			goodsInTrain[name] += amount
		}
		if strings.HasPrefix(line, "delete") {
			trimmed := strings.TrimPrefix(line, "delete ")
			number, _ := strconv.Atoi(trimmed)
			for number != 0 {
				switch {
				case opsNums.Back() > number:
					tmp := opsNums.Pop()
					opsNums.Push(tmp - number)
					goodsInTrain[opsNames.Back()] -= number
					number = 0
				case opsNums.Back() == number:
					tmp := opsNums.Pop()
					tmpName := opsNames.Pop()
					goodsInTrain[tmpName] -= tmp
					number = 0
				case opsNums.Back() < number:
					tmp := opsNums.Pop()
					tmpName := opsNames.Pop()
					goodsInTrain[tmpName] -= tmp
					number -= tmp
				}
			}
		}
	}
}
func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
func readShortLine(r *bufio.Reader) string {
	line, _, _ := r.ReadLine()
	return string(line)
}

type StackStr struct {
	stack []string
	len   int
}

func (s *StackStr) Push(n string) {
	s.stack = append(s.stack, n)
	s.len++
}
func (s *StackStr) Pop() string {
	if s.len != 0 {
		result := s.stack[len(s.stack)-1]
		s.len--
		s.stack = s.stack[:s.len]
		return result
	} else {
		panic("Stack is empty")
	}
}
func (s *StackStr) Back() string {
	if s.len != 0 {
		return s.stack[len(s.stack)-1]
	} else {
		panic("Stack is empty")
	}
}

type StackInt struct {
	stack []int
	len   int
}

func (s *StackInt) Push(n int) {
	s.stack = append(s.stack, n)
	s.len++
}
func (s *StackInt) Pop() int {
	if s.len != 0 {
		result := s.stack[len(s.stack)-1]
		s.len--
		s.stack = s.stack[:s.len]
		return result
	} else {
		panic("Stack is empty")
	}
}
func (s *StackInt) Back() int {
	if s.len != 0 {
		return s.stack[len(s.stack)-1]
	} else {
		panic("Stack is empty")
	}
}
