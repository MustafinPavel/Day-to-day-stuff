package main

import "fmt"

func mainTest() {
	array := []int{6, 10, 5, 8, 9, 8, 2, 3, 10, 7}
	fmt.Println("Before: ", array)
	a := mSort(array)
	fmt.Println("After: ", a)
}

// Сортировка методом вставки
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

// Быстрая Сортировка
func quickSortStart(arr []int) []int {
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

// Сортировка Слиянием:
func mSort(m []int) []int {
	if len(m) <= 1 {
		return m
	}

	mid := len(m) / 2
	left := m[:mid]
	right := m[mid:]

	left = mSort(left)
	right = mSort(right)

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
