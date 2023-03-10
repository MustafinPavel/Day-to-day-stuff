package usefulFuncs

import "fmt"

// Сортировка методом вставки / Insertion sort
func InsertionSort(array []int) {
	for i := 0; i < len(array); i++ {
		temp := array[i]
		j := i
		for j > 0 && array[j-1] >= temp {
			array[j] = array[j-1]
			j--
		}
		array[j] = temp
	}
}

// Быстрая Сортировка / Quicksort
func QuickSortStart(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}
func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}
func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

// Сортировка Слиянием / MergeSort
func MergeSort(m []int) []int {
	if len(m) <= 1 {
		return m
	}

	mid := len(m) / 2
	left := m[:mid]
	right := m[mid:]

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(left, right []int) []int {
	var result []int
	for len(left) > 0 || len(right) > 0 {
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				result = append(result, left[0])
				fmt.Printf("Sorting:\t%v\n", result)
				left = left[1:]
			} else {
				result = append(result, right[0])
				fmt.Printf("Sorting:\t%v\n", result)
				right = right[1:]
			}
		} else if len(left) > 0 {
			result = append(result, left[0])
			fmt.Printf("Sorting:\t%v\n", result)
			left = left[1:]
		} else if len(right) > 0 {
			result = append(result, right[0])
			fmt.Printf("Sorting:\t%v\n", result)
			right = right[1:]
		}
	}
	return result
}

// Пирамидальная сортировка (HeapSort)
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
