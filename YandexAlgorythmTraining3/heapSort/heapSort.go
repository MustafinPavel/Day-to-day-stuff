package heapSort

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Пирамидальная сортировка
func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	readSingleInt(in) //skip line
	array := readLongIntSlice(in)
	//logic
	myHeap := Heap{heap: array}
	myHeap.Heapify()
	myHeap.HeapSort()
	//output
	for _, v := range myHeap.heap {
		out.WriteString(strconv.Itoa(v) + " ")
	}
}

type Heap struct {
	heap []int
}

func (h *Heap) HeapSort() {
	l := len(h.heap)
	for i := 0; i < len(h.heap)-2; i++ {
		swap(0, l-1, h.heap)
		pos := 0
		for pos*2+2 < l-1 {
			maxSonIdx := pos*2 + 1
			if h.heap[pos*2+2] > h.heap[pos*2+1] {
				maxSonIdx = pos*2 + 2
			}
			if h.heap[pos] < h.heap[maxSonIdx] {
				swap(pos, maxSonIdx, h.heap)
				pos = maxSonIdx
			} else {
				break
			}
		}
		l -= 1
	}
	if l == 2 && h.heap[0] > h.heap[1] {
		swap(0, 1, h.heap)
	}
}

func (h *Heap) Heapify() {
	for i := len(h.heap) / 2; i >= 0; i-- {
		pos := i
		for pos*2+1 < len(h.heap) {
			maxSonIdx := pos*2 + 1
			if len(h.heap) == pos*2+2 {
				maxSonIdx = pos*2 + 1
			} else if h.heap[pos*2+2] > h.heap[pos*2+1] {
				maxSonIdx = pos*2 + 2
			}
			if h.heap[pos] < h.heap[maxSonIdx] {
				swap(pos, maxSonIdx, h.heap)
				pos = maxSonIdx
			} else {
				break
			}
		}
	}
}

func swap(a, b int, h []int) {
	h[a], h[b] = h[b], h[a]
}

func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
func readLongIntSlice(in *bufio.Reader) []int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
	result := make([]int, 0, 100000)
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

//Нашлась логическая ошибка в Heapify. StressTest() помог её обнаружить.
//Оказалось, что не обрабатывались частные случаи, когда в исходном массиве чётное число
//элементов, и при этом максимальное значение всего массива стоит в самом конце

// func StressTest() {
// 	stressTest := true
// 	for stressTest {
// 		sl := GenerateRandomSlice()
// 		sl2 := make([]int, len(sl))
// 		res := make([]int, len(sl))
// 		copy(sl2, sl)
// 		copy(res, sl)
// 		myHeap2 := Heap{heap: sl}
// 		myHeap2.Heapify()
// 		myHeap2.HeapSort()
// 		sort.Sort(sort.IntSlice(sl2))
// 		if reflect.DeepEqual(sl, sl2) {
// 			fmt.Println("PASSED")
// 		} else {
// 			fmt.Printf("FAILED!\nsrc:\t%v\nres1:\t%v\nres2:\t%v\n", res, sl, sl2)
// 			stressTest = false
// 		}
// 	}
// }
// func GenerateRandomSlice() []int {
// 	res := make([]int, 0, 10)
// 	for i := 0; i < 10; i++ {
// 		res = append(res, rand.Intn(1000000001))
// 	}
// 	return res
// }
