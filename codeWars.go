package main

/*
func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0 {
		return []float64{}
	}
	var result []float64 = make([]float64, 3, 6)
	copy(result, signature[:])
	if n < 3 {
		return signature[:n]
	} else {
		for i := 0; i < n-3; i++ {
			plus := result[i] + result[i+1] + result[i+2]
			result = append(result, plus)
		}
	}
	return result
}

import "strings"
func DuplicateEncode(word string) string {
	temp := make(map[rune]int)
	word = strings.ToLower(word)
	var result string
	for _, v := range word {
		if _, ok := temp[v]; ok {
			temp[v] += 1
		} else {
			temp[v] = 1
		}
	}
	for _, v := range word {
		if temp[v] == 1 {
			result += "("
		} else {
			result += ")"
		}
	}
	return result
}

func HighestRank(nums []int) int {
	temp := make(map[int]int)
	var temp2 []int
	for _, v := range nums {
		if _, ok := temp[v]; ok {
			temp[v] += 1
		} else {
			temp[v] = 1
		}
	}
	var max int
	for _, v := range temp {
		if v >= max {
			max = v
		}
	}
	for k, v := range temp {
		if v == max {
			temp2 = append(temp2, k)
		}
	}
	return FindMax(temp2)
}
func FindMax(nums []int) int {
	var max int
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

import "strconv"
func FakeBin(x string) string {
	var result string
	for _, v := range x {
		n, _ := strconv.Atoi(string(v))
		if n < 5 {
			result += "0"
		} else {
			result += "1"
		}
	}
	return result
}

import "strings"
func AbbrevName(name string) string {
	sl := strings.Fields(name)
	return strings.ToUpper(string(sl[0][0])) + "." + strings.ToUpper(string(sl[1][0]))
}

import "strconv"
func DigitalRoot(n int) int {
AGAIN:
	var allNumbers []int
	strNum := strconv.Itoa(n)
	for _, v := range strNum {
		b, _ := strconv.Atoi(string(v))
		allNumbers = append(allNumbers, b)
	}
	n = 0
	for _, v := range allNumbers {
		n += v
	}
	if n > 9 {
		goto AGAIN
	}
	return n
}

func CountPositivesSumNegatives(numbers []int) []int {
	var result []int = make([]int, 2, 4)
	for _, v := range numbers {
		switch {
		case v > 0:
			result[0]++
		case v < 0:
			result[1] += v
		}
	}
	if result[0] == 0 && result[1] == 0 {
		return nil
	}
	return result
}

import "strings"
func Accum(s string) string {
	var result string
	for k, v := range s {
		result += strings.ToUpper(string(v))
		for i := 0; i < k; i++ {
			result += strings.ToLower(string(v))
		}
		if k < len(s)-1 {
			result += "-"
		}
	}
	return result
}

import "math"
func SevenWondersScience(compasses int, gears int, tablets int) int {
	var result float64
	mr := math.Min(float64(compasses), math.Min(float64(gears), float64(tablets))) * 7
	result = mr + math.Pow(float64(compasses), 2) + math.Pow(float64(gears), 2) + math.Pow(float64(tablets), 2)
	return int(result)
}

//Вся эта функция заменяется одним strings.Fields()
func StringToArray(str string) []string {
	result := []string{}
	var temp string
	for k, v := range str {
		switch {
		case v == 32:
			result = append(result, temp)
			temp = ""
		case k == len(str)-1:
			temp += string(v)
			result = append(result, temp)
			break
		default:
			temp += string(v)
		}
	}
	return result
}
*/
