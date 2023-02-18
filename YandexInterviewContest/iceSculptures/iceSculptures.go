package iceSculptures

import (
	"bufio"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	sculptsAmount, ice, minutesLeft := readLineWithThreeInt64(in)
	sculptWeights := readIntSlice(in)
	//logic
	//make a sorted slice, where [0] is num of sculpt, [1] is ice difference from ideal
	sculpts := fillAndSortSculpts(sculptWeights, sculptsAmount, ice)
	//choose idealSculpts from a slice:
	resultNumber := 0
	idealSculpts := make([]int, 0, 20000)
	if sculptsAmount > 0 {
		for i := 0; ; i++ {
			minutesLeft -= int64(sculpts[i][1])
			if minutesLeft >= 0 {
				resultNumber++
				idealSculpts = append(idealSculpts, sculpts[i][0])
			}
			if i == len(sculpts)-1 || minutesLeft < 0 {
				break
			}

		}
	}
	//output
	out.WriteString(strconv.Itoa(resultNumber) + "\n")
	var sb strings.Builder
	for i := 0; i < len(idealSculpts); i++ {
		sb.WriteString(strconv.FormatInt(int64(idealSculpts[i]), 10) + " ")
	}
	out.WriteString(sb.String())
}

func fillAndSortSculpts(weights []int, amount, idealIce int) [][]int {
	var sculpts [][]int = make([][]int, 0, 20000)
	for i := 0; i < amount; i++ {
		iceDifference := math.Abs(float64(weights[i] - idealIce))
		sculpts = append(sculpts, []int{i + 1, int(iceDifference)})
	}
	//sort by iceDifference
	sort.Slice(sculpts, func(i, j int) bool {
		return sculpts[i][1] < sculpts[j][1]
	})
	return sculpts
}

func readIntSlice(in *bufio.Reader) []int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 300000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	line := string(tmpByteSlice)
	result := make([]int, 0, len(line)/2)
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}

func readLineWithThreeInt64(r *bufio.Reader) (num, ice int, mins int64) {
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	num, _ = strconv.Atoi(slice[0])
	ice, _ = strconv.Atoi(slice[1])
	mins, _ = strconv.ParseInt(slice[2], 10, 64)
	return
}
