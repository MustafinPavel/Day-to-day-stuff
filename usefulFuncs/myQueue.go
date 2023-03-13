package usefulFuncs

type Queue struct {
	queue []int
	head  int
	len   int
}

func (q *Queue) push(n int) {
	q.queue = append(q.queue, n)
	q.len++
}
func (q *Queue) pop() int {
	var result int
	if q.len != 0 {
		q.len--
		result = q.queue[q.head]
		q.head++
	} else {
		panic("empty queue")
	}
	if q.head > 1000 {
		tmp := make([]int, 0, q.len)
		for i := 0; i < q.len; i++ {
			tmp = append(tmp, q.queue[q.head+i])
		}
		q.queue = tmp
		q.head = 0
	}
	return result
}
func (q *Queue) front() int {
	if q.len != 0 {
		return q.queue[q.head]
	} else {
		panic("empty queue")
	}
}
func (q *Queue) size() int {
	return q.len
}
func (q *Queue) clear() {
	q.queue = make([]int, 0, 1000)
	q.len = 0
	q.head = 0
}
