package usefulFuncs

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Считать короткую строку
func readShortLine(r *bufio.Reader) string {
	line, _, _ := r.ReadLine()
	return string(line)
}

// Считать длинную строку
func readLongLine(in *bufio.Reader) string {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 10000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	return string(tmpByteSlice)
}

// Считывает и парсит строку с одним int
func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}

// Считывает короткую строку с набором int и парсит в []int
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

// Считывает длинную строку и делает []string
func readLongSlice(in *bufio.Reader) []string {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 10000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	slice := strings.Fields(string(tmpByteSlice))
	return slice
}

// Считывает длинную строку из чисел и делает []int
func readLongIntSlice(in *bufio.Reader) []int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 10000)
	result := make([]int, 0, 10000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	slice := strings.Fields(string(tmpByteSlice))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}

// Алгоритм бинарного поиска
func StartSearch(tar int, s []int) int {
	return binarySearch(tar, 0, len(s)-1, s)
}
func binarySearch(target int, first int, last int, s []int) int {
	m := math.Round(float64(first + (last-first)/2))
	middle := int(m)
	if first > last {
		panic("no such value in a collection")
	}
	if s[middle] > target {
		return binarySearch(target, first, middle-1, s)
	}
	if s[middle] < target {
		return binarySearch(target, middle+1, last, s)
	}
	return middle
}

// Merge 2 SORTED slices a and b into one
func mergeSlices(a []int, b []int) []int {
	var result []int
	var ai, bi int
Loop:
	for {
		switch {
		case ai == len(a) && bi == len(b):
			break Loop
		case ai == len(a):
			result = append(result, b[bi])
			bi++
		case bi == len(a):
			result = append(result, a[ai])
			ai++
		case a[ai] <= b[bi]:
			result = append(result, a[ai])
			ai++
		case b[bi] < a[ai]:
			result = append(result, b[bi])
			bi++
		}
	}
	return result
}

// Reversing a slice with O(1) memory:
func reverseSlice(b []int) {
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
	}
}

// Проверка того, имплементировала ли структура ВСЕ методы интерфейса.
// Если добавить метод в интерфейс, не даст скомпилироваться
type SomeInterface interface {
	Method()
}
type Implementation struct{}

func (*Implementation) Method() { fmt.Println("Hello, World!") }

var _ SomeInterface = (*Implementation)(nil) // ← вот эта строчка

// Поиск NOD (Наименьший общий делитель)
// и NOK (Наименьшее общее кратное)
func findNOD(a, b int) int {
	for {
		max := int(math.Max(float64(a), float64(b)))
		min := int(math.Min(float64(a), float64(b)))
		if max%min == 0 {
			return min
		} else {
			result := findNOD(min, max%min)
			return result
		}
	}
}
func findNOK(a, b int) int {
	nod := findNOD(a, b)
	return a * b / nod
}

//Проверка выделенной памяти
// var m runtime.MemStats
// runtime.ReadMemStats(&m)
// fmt.Println(m.TotalAlloc/1024/1024, " Mb")
