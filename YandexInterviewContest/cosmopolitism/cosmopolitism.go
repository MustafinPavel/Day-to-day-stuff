package cosmopolitism

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Я очень долго думал, как оптимизировать этот наивный алгоритм из О(Q^N),
// но так ничего и не придумал, кроме как создать разные слайсы стран
// для людей с разными возможностями, чтобы сразу отсекать половину вариантов для
// людей без образования, например, или, отсортировав массив по необходимому доходу,
// сразу отсекать всех со слишком высокими требованиями.
// Возможно и это не тот путь...надо подумать -_-
func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	amountOfCountries := readLineWithOneInt(in)
	reqIncome := readIntSlice(in)
	reqEducation := readIntSlice(in)
	reqCitizenship := readIntSlice(in)
	amountOfClassmates := readLineWithOneInt(in)
	clsIncome := readIntSlice(in)
	clsEducation := readIntSlice(in)
	clsCitizenship := readIntSlice(in)
	//logic
	for i := 0; i < amountOfClassmates; i++ {
		countryNotFound := true
		for j := 0; countryNotFound && j < amountOfCountries; j++ {
			if (hasEducation(i, j, reqEducation, clsEducation) && incomeIsEnough(i, j, reqIncome, clsIncome)) || parentsHasCitizenship(i, j, reqCitizenship, clsCitizenship) {
				out.WriteString(strconv.Itoa(j+1) + " ")
				countryNotFound = false
			} else if j == amountOfCountries-1 {
				out.WriteString("0 ")
			}
		}
	}
}

func parentsHasCitizenship(mateNum int, countryNum int, reqCitizenship []int, prntCitiznshp []int) bool {
	if prntCitiznshp[mateNum]-1 == countryNum && reqCitizenship[countryNum] == 1 {
		return true
	}
	return false
}

func incomeIsEnough(mateNum int, countryNum int, reqIncome []int, clsIncome []int) bool {
	return reqIncome[countryNum] <= clsIncome[mateNum]
}

func hasEducation(mateNum int, countryNum int, reqEduc []int, clsEduc []int) bool {
	if reqEduc[countryNum] == 0 || (reqEduc[countryNum] == 1 && clsEduc[mateNum] == 1) {
		return true
	}
	return false
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

func readLineWithOneInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
