package usefulFuncs

type Heap struct {
	heap []int
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
func (h *Heap) Push(n int) {
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

func (h *Heap) Pop() (res int) {
	res = h.heap[0]
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
	return
}

func swap(a, b int, h []int) {
	h[a], h[b] = h[b], h[a]
}
