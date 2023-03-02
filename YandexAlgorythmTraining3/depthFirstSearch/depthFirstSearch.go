package depthFirstSearch

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
	V := info[0] //1000
	E := info[1] //50000
	//логика
	graph := make([][]int, V+1, 1001)
	visited := make([]bool, V+1, 1001)
	for i := 0; i < E; i++ {
		tmp := readShortIntSlice(in)
		if tmp[0] == tmp[1] {
			continue
		}
		graph[tmp[0]] = append(graph[tmp[0]], tmp[1])
		graph[tmp[1]] = append(graph[tmp[1]], tmp[0])
	}
	dfs(graph, visited, 1)

	//вывод
	var result string
	var counter int
	for i := 0; i < len(visited); i++ {
		if visited[i] {
			counter++
			result += " " + strconv.Itoa(i)
		}
	}
	result = strings.TrimSpace(result)
	out.WriteString(strconv.Itoa(counter) + "\n" + result)

}

func dfs(graph [][]int, visited []bool, now int) int {
	visited[now] = true
	for _, neigh := range graph[now] {
		if !visited[neigh] {
			dfs(graph, visited, neigh)
		}
	}
	return 0
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
