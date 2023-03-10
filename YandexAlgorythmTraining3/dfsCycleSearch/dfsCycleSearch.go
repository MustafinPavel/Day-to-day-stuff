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
	V := readSingleInt(in) //500
	graphTable := make([][]int, 0, V)
	for i := 0; i < V; i++ {
		tmp := readShortIntSlice(in)
		graphTable = append(graphTable, tmp)
	}
	//логика
	graph := make([][]int, V+1, V+1)
	visited := make([]int, V+1, V+1)
	for i := 0; i < len(graphTable); i++ {
		for j := 0; j < len(graphTable[i]); j++ {
			if graphTable[i][j] == 1 {
				graph[i+1] = append(graph[i+1], j+1)
			}
		}
	}

	var st Stack
	var r int
	for i := 1; i < len(visited); i++ {
		if visited[i] == 0 {
			st, r, _ = dfs(graph, visited, i, 0, false, &st)
		}
	}
	//вывод
	if r == 0 {
		out.WriteString("NO")
	} else {
		out.WriteString("YES\n" + strconv.Itoa(len(st.stack)) + "\n")
		itr := len(st.stack)
		for i := 0; i < itr; i++ {
			out.WriteString(strconv.Itoa(st.Pop()))
			if i != itr-1 {
				out.WriteString(" ")
			}
		}
		//
	}
}

// Квинтэссенция говногода
func dfs(graph [][]int, visited []int, now int, prev int, cycleFound bool, st *Stack) (result Stack, r int, cf bool) {
	visited[now] = 1
	st.Push(now)
	for _, neigh := range graph[now] {
		if cycleFound {
			return *st, r, cycleFound
		}
		if visited[neigh] == 1 && neigh != prev {
			cycleFound = true
			r = neigh
			return *st, r, cycleFound
		}
		if visited[neigh] == 0 {
			prev = now
			result, r, cycleFound = dfs(graph, visited, neigh, prev, cycleFound, st)
		}
	}
	if !cycleFound {
		visited[now] = 2
		st.Pop()
	}
	return
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
