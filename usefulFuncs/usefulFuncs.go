package usefulFuncs

import (
	"bufio"
	"fmt"
	"math"
	"os"
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
	tmpByteSlice := make([]byte, 0, 100000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	return string(tmpByteSlice)
}

// Считывает и парсит строку с одним int
func readLineWithOneInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}

// Считывает короткую строку с набором int и парсит в []int
func sliceFromShortLine(r *bufio.Reader) []int {
	result := make([]int, 0, 100000)
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
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

// Stack data structure (LIFO) - Lazy one:
func DataStructures() {
	var s []int = []int{1, 2, 3, 4, 5}
	s, lifo := s[:len(s)-1], s[len(s)-1]
	fmt.Println(s, lifo)
}

// Stack data structure (LIFO):
type Stack struct {
	stack []string
}

func (st *Stack) AddToStack(s string) {
	st.stack = append(st.stack, s)
}
func (st *Stack) PopFromStack() string {
	if len(st.stack) != 0 {
		returningString := st.stack[len(st.stack)-1]
		st.stack = st.stack[:len(st.stack)-1]
		return returningString
	}
	panic("The stack is empty")
}

// Reversing a slice with O(1) memory:
func reverseSlice(b []int) {
	for i := 0; i < len(b)/2; i++ {
		b[i] = b[len(b)-1-i] * b[i]
		b[len(b)-1-i] = b[i] / b[len(b)-1-i]
		b[i] = b[i] / b[len(b)-1-i]
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

// Генерация строки с i целыми числами и запись её в output.txt
// для больших тестов решений задач
func RandNumSliceGen() {
	file1, _ := os.Create("output.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(file1)
	for i := 0; i < 10000; i++ {
		out.WriteString(strconv.Itoa(88) + " ")
	}
	out.Flush()
	isNotEnded := true
	parts := 0
	numSlice := make([]int, 0, 10000)
	for isNotEnded {
		parts++
		cds, end := readIntSlice(in)
		isNotEnded = end
		numSlice = append(numSlice, cds...)
	}
	fmt.Printf("num slice len:\t%v\nparts:\t%v\n", len(numSlice), parts)
}
func readIntSlice(in *bufio.Reader) ([]int, bool) {
	temp, isNotEnded, _ := in.ReadLine()
	tempSl := strings.Split(string(temp), " ")
	var result []int = make([]int, 0, len(tempSl)*2)
	for _, v := range tempSl {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}
	return result, isNotEnded
}
