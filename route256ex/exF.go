package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func exF() {
	// file1, _ := os.Open("testFileInput")
	// r := bufio.NewReader(file1)
	r := bufio.NewReader(os.Stdin)
	testsAmount := ReadIntFromLine(r)
	for i := 0; i < testsAmount; i++ {
		localInputsAmount := ReadIntFromLine(r)
		fmt.Println(ScanAndCount(localInputsAmount, r))
	}
}

func ScanAndCount(linesAmount int, r *bufio.Reader) string {
	var pairs [][]int = make([][]int, 0, 10000)
	for i := 0; i < linesAmount; i++ {
		line, _, _ := r.ReadLine()
		timePair := strings.Split(string(line), "-")
		timePairInt, ok := TurnTimeIntoIntAndCheck(timePair)
		if !ok {
			for j := i + 1; j < linesAmount; j++ {
				r.ReadLine()
			}
			return "NO"
		}
		pairs = append(pairs, timePairInt)
	}
	ok := CheckForIntersections(pairs)
	if !ok {
		return "NO"
	}
	return "YES"
}

func TurnTimeIntoIntAndCheck(pair []string) ([]int, bool) {
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

func CheckForIntersections(pairs [][]int) bool {
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

func ReadIntFromLine(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
