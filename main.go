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
	info := readShortIntSlice(in)
	V := info[0] //100
	E := info[1] //4450
	//логика
	graph := make([][]int, V+1, V+1)
	visited := make([]int, V+1, V+1)
	for i := 0; i < E; i++ {
		tmp := readShortIntSlice(in)
		if tmp[0] == tmp[1] {
			continue
		}
		graph[tmp[0]] = append(graph[tmp[0]], tmp[1])
		graph[tmp[1]] = append(graph[tmp[1]], tmp[0])
	}

	isDividable := true
	for i := 1; i < len(visited) && isDividable; i++ {
		if visited[i] == 0 {
			isDividable = dfs(graph, visited, i, 1, isDividable)
		}
	}
	//вывод
	if isDividable {
		out.WriteString("YES")
	} else {
		out.WriteString("NO")
	}
}

func dfs(graph [][]int, visited []int, now int, color int, dividable bool) bool {
	visited[now] = color
	for _, neigh := range graph[now] {
		if visited[neigh] == color {
			dividable = false
		}
		if visited[neigh] == 0 {
			dividable = dfs(graph, visited, neigh, 3-color, dividable)
		}
	}
	return dividable
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
