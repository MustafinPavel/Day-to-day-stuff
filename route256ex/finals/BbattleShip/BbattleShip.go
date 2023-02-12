package BbattleShip

import (
	"bufio"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

// Проверяет корректность состава кораблей морского боя.
func battleShipCheck() {
	r := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var shipsPattern sort.IntSlice = []int{1, 1, 1, 1, 2, 2, 2, 3, 3, 4}
	numOfStrings := readIntFromLine(r)
	for i := 0; i < numOfStrings; i++ {
		var stringWithNums sort.IntSlice = readSliceFromLine(r)
		sort.Sort(stringWithNums)
		if reflect.DeepEqual(stringWithNums, shipsPattern) {
			out.WriteString("YES\n")
		} else {
			out.WriteString("NO\n")
		}
	}
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
