package usefulFuncs

import (
	"fmt"
	"os"
)

type Queue struct {
	queue []int
	head  int
	len   int
}

func (q *Queue) push(n int) {
	q.queue = append(q.queue, n)
	q.len++
	fmt.Println("ok")
}
func (q *Queue) pop() {
	if q.len != 0 {
		fmt.Printf("%v\n", q.queue[q.head])
		q.len--
		q.head++
	} else {
		fmt.Println("error")
	}
	if q.head > 1000 {
		tmp := make([]int, 0, q.len)
		for i := 0; i < q.len; i++ {
			tmp = append(tmp, q.queue[q.head+i])
		}
		q.queue = tmp
		q.head = 0
	}
}
func (q *Queue) front() {
	if q.len != 0 {
		fmt.Printf("%v\n", q.queue[q.head])
	} else {
		fmt.Println("error")
	}
}
func (q *Queue) size() {
	fmt.Printf("%v\n", q.len)
}
func (q *Queue) clear() {
	q.queue = make([]int, 0, 1000)
	q.len = 0
	q.head = 0
	fmt.Println("ok")
}
func (q *Queue) exit() {
	fmt.Println("bye")
	os.Exit(0)
}
