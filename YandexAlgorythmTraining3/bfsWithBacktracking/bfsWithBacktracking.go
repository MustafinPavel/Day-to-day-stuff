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
	graphSize := readSingleInt(in)
	graph := make([][]int, graphSize+1, graphSize+1)
	for i := 0; i < graphSize; i++ {
		line := readShortIntSlice(in)
		for j := 0; j < len(line); j++ {
			if line[j] == 1 {
				graph[i+1] = append(graph[i+1], j+1)
			}
		}
	}
	tmp := readShortIntSlice(in)
	start := tmp[0]
	end := tmp[1]
	//логика
	var queue Queue
	results := make([]int, graphSize+1, graphSize+1)
	wayback := make([]int, graphSize+1, graphSize+1)
	queue.push(start)
BFS:
	for queue.len > 0 {
		now := queue.pop()
		for i := 0; i < len(graph[now]); i++ {
			neigh := graph[now][i]
			if results[neigh] == 0 && neigh != start {
				queue.push(neigh)
				results[neigh] = results[now] + 1
				wayback[neigh] = now
			}
			if neigh == end {
				break BFS
			}
		}
	}
	//вывод
	if results[end] == 0 {
		switch {
		case start == end:
			out.WriteString("0")
		case start != end:
			out.WriteString("-1")
		}
	} else {
		way := make([]int, 0)
		out.WriteString(strconv.Itoa(results[end]) + "\n")
		now := end
		for wayback[now] != 0 {
			way = append(way, now)
			now = wayback[now]
		}
		way = append(way, now)
		for i := len(way) - 1; i >= 0; i-- {
			out.WriteString(strconv.Itoa(way[i]))
			if i != 0 {
				out.WriteString(" ")
			}
		}
	}
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

type Queue struct {
	queue []int
	head  int
	len   int
}

func (q *Queue) push(n int) {
	q.queue = append(q.queue, n)
	q.len++
}
func (q *Queue) pop() int {
	var result int
	if q.len != 0 {
		q.len--
		result = q.queue[q.head]
		q.head++
	} else {
		panic("empty queue")
	}
	if q.head > 1000 {
		tmp := make([]int, 0, q.len)
		for i := 0; i < q.len; i++ {
			tmp = append(tmp, q.queue[q.head+i])
		}
		q.queue = tmp
		q.head = 0
	}
	return result
}
