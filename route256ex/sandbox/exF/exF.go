package exF

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Не очень изящное, но неплохо структурированное решение
func Solution() {
	// file1, _ := os.Open("testFileInput")
	// r := bufio.NewReader(file1)
	r := bufio.NewReader(os.Stdin)
	testsAmount := readIntFromLine(r)
	for i := 0; i < testsAmount; i++ {
		localInputsAmount := readIntFromLine(r)
		fmt.Println(scanAndCount(localInputsAmount, r))
	}
}

func scanAndCount(linesAmount int, r *bufio.Reader) string {
	var pairs [][]int = make([][]int, 0, 10000)
	for i := 0; i < linesAmount; i++ {
		line, _, _ := r.ReadLine()
		timePair := strings.Split(string(line), "-")
		timePairInt, ok := turnTimeIntoIntAndCheck(timePair)
		if !ok {
			for j := i + 1; j < linesAmount; j++ {
				r.ReadLine()
			}
			return "NO"
		}
		pairs = append(pairs, timePairInt)
	}
	ok := checkForIntersections(pairs)
	if !ok {
		return "NO"
	}
	return "YES"
}

func turnTimeIntoIntAndCheck(pair []string) ([]int, bool) {
	var res []int = make([]int, 0, 1000)
	for _, v := range pair {
		ts := strings.Split(v, ":")
		var tsi []int
		for i := 0; i < len(ts); i++ {
			t, _ := strconv.Atoi(ts[i])
			tsi = append(tsi, t)
		}
		if tsi[0] > 23 || tsi[1] > 59 || tsi[2] > 59 {
			return nil, false
		}
		res = append(res, tsi[2]+tsi[1]*60+tsi[0]*3600)
	}
	if res[0] > res[1] {
		return nil, false
	}
	return res, true
}

func checkForIntersections(pairs [][]int) bool {
	for k, v := range pairs {
		for i := 0; i < len(pairs); i++ {
			switch {
			case k == i:
				continue
			case v[0] >= pairs[i][0] && v[0] <= pairs[i][1]:
				return false
			case v[1] >= pairs[i][0] && v[1] <= pairs[i][1]:
				return false
			}
		}
	}
	return true
}

func readIntFromLine(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
