// Принята не полностью
package DraceResults

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Нумерует места спортсменов. Если разница меньше 1cек, то считает за одно место.
func OrderSportsmenByPlaces() {
	r := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	runsAmount := readIntFromLine(r)
	for i := 0; i < runsAmount; i++ {
		numOfSportsman := readIntFromLine(r)
		sportsmanSlice := readSliceFromLine(r)
		places := getPlaces(numOfSportsman, sportsmanSlice)
		for _, v := range sportsmanSlice {
			out.WriteString(strconv.Itoa(places[v]) + " ")
		}
		out.WriteString("\n")
	}
}

func getPlaces(numOfSportsman int, sportsmanSlice []int) map[int]int {
	resultsMap := make(map[int]int)
	t := make([]int, len(sportsmanSlice))
	copy(t, sportsmanSlice)
	var temp sort.IntSlice = t
	sort.Sort(temp)
	for i := 0; i < len(temp); i++ {
		if i == 0 {
			resultsMap[temp[i]] = 1
			continue
		}
		if temp[i]-temp[i-1] <= 1 {
			resultsMap[temp[i]] = resultsMap[temp[i-1]]
		} else {
			resultsMap[temp[i]] = i + 1
		}
	}
	return resultsMap
}

func readIntFromLine(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}

func readSliceFromLine(r *bufio.Reader) []int {
	result := make([]int, 0, 100000)
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}
