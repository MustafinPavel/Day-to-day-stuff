package heapUI

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	reqAmount := readSingleInt(in)
	var myHeap Heap
	for i := 0; i < reqAmount; i++ {
		text := readShortLine(in)
		handleQuery(text, &myHeap)
	}
}

type Heap struct {
	heap []int
}

func (h *Heap) Insert(n int) {
	h.heap = append(h.heap, n)
	pos := len(h.heap) - 1
	for (pos-1)/2 >= 0 {
		fatherIdx := (pos - 1) / 2
		if h.heap[fatherIdx] < h.heap[pos] {
			swap(pos, fatherIdx, h.heap)
			pos = fatherIdx
		} else {
			break
		}
	}
}

func (h *Heap) Extract() {
	res := h.heap[0]
	h.heap[0] = h.heap[len(h.heap)-1]
	pos := 0
	for pos*2+2 < len(h.heap) {
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
	h.heap = h.heap[:len(h.heap)-1]
	fmt.Println(res)
}

func (h *Heap) Heapify() {
	for i := len(h.heap) / 2; i >= 0; i-- {
		pos := i
		for pos*2+2 < len(h.heap) {
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
	}
}

func swap(a, b int, h []int) {
	h[a], h[b] = h[b], h[a]
}

func handleQuery(query string, h *Heap) {
	switch {
	case strings.HasPrefix(query, "0 "):
		query = strings.TrimPrefix(query, "0 ")
		n, _ := strconv.Atoi(query)
		h.Insert(n)
	case strings.HasPrefix(query, "1"):
		h.Extract()
	case strings.HasPrefix(query, "EOF"):
		os.Exit(1)
	}
}

func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
func readShortLine(r *bufio.Reader) string {
	line, _, _ := r.ReadLine()
	return string(line)
}
