package drunCardgame

import (
	"bufio"
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
	line1 := readSliceFromShortLine(in)
	line2 := readSliceFromShortLine(in)
	//logic
	player1 := Queue{queue: line1, head: 0, len: len(line1)}
	player2 := Queue{queue: line2, head: 0, len: len(line2)}
	round, roundLimit := 0, 1000000
	for player1.len > 0 && player2.len > 0 {
		p1card, p2card := player1.pop(), player2.pop()
		winner := chooseWinner(p1card, p2card, &player1, &player2)
		winner.push(p1card)
		winner.push(p2card)
		round++
		if round >= roundLimit {
			out.WriteString("botva")
			break
		}
		if player1.len == 0 {
			out.WriteString("second " + strconv.Itoa(round))
		}
		if player2.len == 0 {
			out.WriteString("first " + strconv.Itoa(round))
		}
	}
}

func chooseWinner(p1card, p2card int, p1, p2 *Queue) *Queue {
	switch {
	case p1card == 0 && p2card == 9:
		return p1
	case p2card == 0 && p1card == 9:
		return p2
	case p1card > p2card:
		return p1
	case p2card > p1card:
		return p2
	}
	panic("someone is cheating")
}

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
	if q.head > 1000 {
		tmp := make([]int, 0, q.len)
		for i := 0; i < q.len; i++ {
			tmp = append(tmp, q.queue[q.head+i])
		}
		q.queue = tmp
		q.head = 0
	}
	if q.len != 0 {
		res := q.queue[q.head]
		q.len--
		q.head++
		return res
	}
	panic("queue is empty")
}

func readSliceFromShortLine(r *bufio.Reader) []int {
	result := make([]int, 0, 100000)
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}
