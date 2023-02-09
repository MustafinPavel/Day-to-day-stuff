package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	value int
	next  *Node
}
type MyLinkedList struct {
	head *Node
	len  int
}

func (n *MyLinkedList) Add(a int) {
	newNode := Node{
		value: a,
		next:  nil,
	}
	if n.len == 0 {
		n.head = &newNode
		n.len++
		return
	}
	ptr := n.head
	for i := 0; i < n.len; i++ {
		if ptr.next == nil {
			ptr.next = &newNode
			n.len++
			return
		}
		ptr = ptr.next
	}
}
func (l *MyLinkedList) PrintList() {
	resultString := ""
	ptr := l.head
	for i := 0; i < l.len; i++ {
		resultString = resultString + strconv.Itoa(ptr.value) + " "
		ptr = ptr.next
	}
	fmt.Println(resultString)
}
func (l *MyLinkedList) Remove(val int) {
	ptr := l.head
	ptr2 := ptr.next
	for i := 0; i < l.len; i++ {
		if i == 0 && ptr.value == val {
			l.head = ptr2
			l.len--
			return
		}
		if ptr2.value == val {
			ptr.next = ptr2.next
			l.len--
			return
		}
		ptr = ptr2
		ptr2 = ptr.next
	}
}
