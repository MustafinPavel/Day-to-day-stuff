package dfsTopologicSort

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
	V := info[0] //100000
	E := info[1] //100000
	//логика
	graph := make([][]int, V+1, V+1)
	visited := make([]int, V+1, V+1)
	for i := 0; i < E; i++ {
		tmp := readShortIntSlice(in)
		graph[tmp[0]] = append(graph[tmp[0]], tmp[1])
	}

	isSortable := true
	result := make([]int, 0, V+1)
	for i := 1; i < len(visited) && isSortable; i++ {
		if visited[i] == 0 {
			isSortable, result = dfs(graph, visited, i, isSortable, result)
		}
	}
	reverse(result)
	//вывод
	if isSortable {
		for i := 0; i < len(result); i++ {
			out.WriteString(strconv.Itoa(result[i]))
			if i != len(result)-1 {
				out.WriteString(" ")
			}
		}
	} else {
		out.WriteString("-1")
	}
}

func reverse(s []int) {
	for i := 0; i < len(s)/2; i++ {
		s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	}
}

func dfs(graph [][]int, visited []int, now int, isSortable bool, result []int) (bool, []int) {
	visited[now] = 1
	for _, neigh := range graph[now] {
		if visited[neigh] == 1 {
			isSortable = false
		}
		if visited[neigh] == 0 {
			isSortable, result = dfs(graph, visited, neigh, isSortable, result)
		}
	}
	visited[now] = 2
	result = append(result, now)
	return isSortable, result
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
