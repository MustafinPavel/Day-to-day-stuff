package dfsComponents

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
		if tmp[0] == tmp[1] {
			continue
		}
		graph[tmp[0]] = append(graph[tmp[0]], tmp[1])
		graph[tmp[1]] = append(graph[tmp[1]], tmp[0])
	}
	var tally int
	for i := 1; i < len(visited); i++ {
		if visited[i] == 0 {
			tally++
			dfs(graph, visited, i, tally)
		}
	}
	//вывод
	out.WriteString(strconv.Itoa(tally) + "\n")
	var resSlice [][]int = make([][]int, tally+1, tally+1)
	for i := 1; i < len(visited); i++ {
		resSlice[visited[i]] = append(resSlice[visited[i]], i)
	}
	for k, v := range resSlice {
		if k > 0 {
			var result string
			out.WriteString(strconv.Itoa(len(v)) + "\n")
			for _, v2 := range v {
				result += strconv.Itoa(v2) + " "
			}
			result = strings.TrimSpace(result)
			out.WriteString(result + "\n")
		}
	}
}

func dfs(graph [][]int, visited []int, now int, tally int) {
	visited[now] = tally
	for _, neigh := range graph[now] {
		if visited[neigh] == 0 {
			dfs(graph, visited, neigh, tally)
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
