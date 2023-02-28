package calculator

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// Решение почему-то не проходит по памяти, хотя замеры через runtime
// показывают расход всего 7Мб. Вариант с memo через map расходует 80+Мб.
func main() {
	//ввод
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	N := readSingleInt(in)
	//логика
	memo := make([]int, 1000001)
	ans := dp(N, memo)
	//вывод
	out.WriteString(strconv.Itoa(ans) + "\n")
	var st Stack
	st.Push(N)
	for N != 1 {
		if N%3 == 0 && memo[N]-memo[N/3] == 1 {
			st.Push(N / 3)
			N = N / 3
		}
		if N%2 == 0 && memo[N]-memo[N/2] == 1 {
			st.Push(N / 2)
			N = N / 2
		}
		if memo[N]-memo[N-1] == 1 {
			st.Push(N - 1)
			N = N - 1
		}
	}
	elems := len(st.stack)
	for i := 0; i < elems; i++ {
		out.WriteString(strconv.Itoa(st.Pop()))
		if i != elems-1 {
			out.WriteString(" ")
		}
	}
}

func dp(N int, memo []int) int {
	if memo[N] != 0 {
		return memo[N]
	}
	if N == 0 {
		return 0
	}
	if N == 1 {
		return 0
	}
	if N > 1 {
		min := dp(N-1, memo) + 1
		if N%3 == 0 {
			min = int(math.Min(float64(min), float64(dp(N/3, memo)+1)))
		}
		if N%2 == 0 {
			min = int(math.Min(float64(min), float64(dp(N/2, memo)+1)))
		}
		memo[N] = min
		return min
	}
	panic("")
}

func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}

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

//Вариант с memo в виде Map:
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"math"
// 	"os"
// 	"runtime"
// 	"strconv"
// )

// func main() {
// 	//ввод
// 	file1, _ := os.Open("input.txt")
// 	in := bufio.NewReader(file1)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()
// 	N := readSingleInt(in)

// 	//логика
// 	memo := make(map[int]int)
// 	ans := dp(N, memo)
// 	var m runtime.MemStats
// 	runtime.ReadMemStats(&m)
// 	fmt.Println(m.TotalAlloc/1024/1024, " Mb")
// 	//вывод
// 	out.WriteString(strconv.Itoa(ans) + "\n")
// 	var st Stack
// 	st.Push(N)
// 	for N != 1 {
// 		if N%3 == 0 && memo[N]-memo[N/3] == 1 {
// 			st.Push(N / 3)
// 			N = N / 3
// 		}
// 		if N%2 == 0 && memo[N]-memo[N/2] == 1 {
// 			st.Push(N / 2)
// 			N = N / 2
// 		}
// 		if memo[N]-memo[N-1] == 1 {
// 			st.Push(N - 1)
// 			N = N - 1
// 		}
// 	}
// 	elems := len(st.stack)
// 	for i := 0; i < elems; i++ {
// 		out.WriteString(strconv.Itoa(st.Pop()))
// 		if i != elems-1 {
// 			out.WriteString(" ")
// 		}
// 	}
// }

// func dp(N int, memo map[int]int) int {
// 	if _, ok := memo[N]; ok {
// 		return memo[N]
// 	}
// 	if N == 0 {
// 		return 0
// 	}
// 	if N == 1 {
// 		return 0
// 	}
// 	if N > 1 {
// 		min := dp(N-1, memo) + 1
// 		if N%3 == 0 {
// 			min = int(math.Min(float64(min), float64(dp(N/3, memo)+1)))
// 		}
// 		if N%2 == 0 {
// 			min = int(math.Min(float64(min), float64(dp(N/2, memo)+1)))
// 		}
// 		memo[N] = min
// 		return min
// 	}
// 	panic("")
// }

// func readSingleInt(r *bufio.Reader) int {
// 	line, _, _ := r.ReadLine()
// 	lineInt, _ := strconv.Atoi(string(line))
// 	return lineInt
// }

// type Stack struct {
// 	stack []int
// 	len   int
// }

// func (s *Stack) Push(n int) {
// 	s.stack = append(s.stack, n)
// 	s.len++
// }

// func (s *Stack) Pop() int {
// 	if s.len != 0 {
// 		result := s.stack[len(s.stack)-1]
// 		s.len--
// 		s.stack = s.stack[:s.len]
// 		return result
// 	} else {
// 		panic("Stack is empty")
// 	}
// }
