package deQueueWithErrProtection

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Дек (double ended queue), реализованный через кольцевой буфер
// Надо будет допилить автоматическое сужение при существенном падении числа элементов.
func main() {
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	myDEQueue := NewDEQueue()
	for {
		text := readShortLine(in)
		handleQuery(text, &myDEQueue)
	}
}

type DEQueue struct {
	dqueue []int
	head   int
	tail   int
	len    int
	boofer int
}

func NewDEQueue() DEQueue {
	return DEQueue{
		boofer: 5,
		dqueue: make([]int, 5),
		head:   0,
		tail:   0,
		len:    0,
	}
}

func (q *DEQueue) push_back(n int) {
	if q.len != 0 {
		q.tail++
	}
	q.len++
	if q.len == q.boofer {
		tmp := make([]int, q.boofer*2)
		for i := 0; i < q.len; i++ {
			tmp[i] = q.dqueue[(q.head+i)%q.boofer]
		}
		q.head = 0
		q.tail = q.len - 1
		q.boofer *= 2
		q.dqueue = tmp
	}
	q.dqueue[q.tail%q.boofer] = n
	fmt.Println("ok")
}

func (q *DEQueue) push_front(n int) {
	if q.len != 0 {
		q.head--
		if q.head < 0 {
			q.head = (q.boofer - 1) % q.boofer
		}
	}
	q.len++
	if q.len == q.boofer {
		tmp := make([]int, q.boofer*2)
		for i := 0; i < q.len; i++ {
			tmp[i] = q.dqueue[(q.head+i)%q.boofer]
		}
		q.head = 0
		q.tail = q.len - 1
		q.boofer *= 2
		q.dqueue = tmp
	}
	q.dqueue[q.head%q.boofer] = n

	fmt.Println("ok")
}

func (q *DEQueue) pop_front() {
	if q.len != 0 {
		fmt.Printf("%v\n", q.dqueue[q.head%q.boofer])
		if q.head%q.boofer != q.tail%q.boofer {
			q.head++
		}
		q.len--
	} else {
		fmt.Println("error")
	}
}

func (q *DEQueue) pop_back() {
	if q.len != 0 {
		fmt.Printf("%v\n", q.dqueue[q.tail%q.boofer])
		if q.head%q.boofer != q.tail%q.boofer {
			q.tail--
			if q.tail < 0 {
				q.tail = (q.boofer - 1) % q.boofer
			}
		}
		q.len--
	} else {
		fmt.Println("error")
	}
}

func (q *DEQueue) front() {
	if q.len != 0 {
		fmt.Printf("%v\n", q.dqueue[q.head%q.boofer])
	} else {
		fmt.Println("error")
	}
}

func (q *DEQueue) back() {
	if q.len != 0 {
		fmt.Printf("%v\n", q.dqueue[q.tail%q.boofer])
	} else {
		fmt.Println("error")
	}
}

func (q *DEQueue) size() {
	fmt.Printf("%v\n", q.len)
}

func (q *DEQueue) clear() {
	q.dqueue = NewDEQueue().dqueue
	q.len = NewDEQueue().len
	q.head = NewDEQueue().head
	q.tail = NewDEQueue().tail
	fmt.Println("ok")
}

func (q *DEQueue) exit() {
	fmt.Println("bye")
	os.Exit(0)
}

func handleQuery(query string, q *DEQueue) {
	switch {
	case strings.HasPrefix(query, "push_back"):
		query = strings.TrimPrefix(query, "push_back ")
		n, _ := strconv.Atoi(query)
		q.push_back(n)
	case strings.HasPrefix(query, "push_front"):
		query = strings.TrimPrefix(query, "push_front ")
		n, _ := strconv.Atoi(query)
		q.push_front(n)
	case strings.HasPrefix(query, "pop_back"):
		q.pop_back()
	case strings.HasPrefix(query, "pop_front"):
		q.pop_front()
	case strings.HasPrefix(query, "front"):
		q.front()
	case strings.HasPrefix(query, "back"):
		q.back()
	case strings.HasPrefix(query, "size"):
		q.size()
	case strings.HasPrefix(query, "clear"):
		q.clear()
	case strings.HasPrefix(query, "exit"):
		q.exit()
	case strings.HasPrefix(query, "EOF"):
		os.Exit(1)
	}
}

func readShortLine(r *bufio.Reader) string {
	line, _, err := r.ReadLine()
	if err != nil {
		return "EOF"
	}
	return string(line)
}
