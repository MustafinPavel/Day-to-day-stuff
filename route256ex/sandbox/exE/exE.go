package exE

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Неплохая структура
func Solution() {
	// file1, _ := os.Open("testFileInput")
	// r := bufio.NewReader(file1)
	r := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	employeesAmount := readIntFromLine(r)
	for i := 0; i < employeesAmount; i++ {
		daysAmount := readIntFromLine(r)
		reports := readSliceFromLine(r)
		fmt.Println(checkEmployee(daysAmount, reports))
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

func checkEmployee(daysAmount int, reports []int) (testResult string) {
	sliceToCompare := make([]int, 0, cap(reports))
	var result bool = true
	for k, v := range reports {
		if k == 0 {
			sliceToCompare = append(sliceToCompare, v)
		} else {
			if v == reports[k-1] {
				continue
			} else {
				result = checkSlice(sliceToCompare, v)
				if !result {
					return "NO"
				}
				sliceToCompare = append(sliceToCompare, v)
			}
		}
	}
	return "YES"
}

func checkSlice(slice []int, v int) (check bool) {
	for i := 0; i < len(slice); i++ {
		if slice[i] == v {
			return false
		}
	}
	return true
}
