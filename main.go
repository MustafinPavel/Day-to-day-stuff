package main

import (
	"fmt"
)

func main() {
	var testStack Stack
	testStack.AddToStack("one")
	testStack.AddToStack("two")
	fmt.Println(testStack)
	fmt.Println(testStack.PopFromStack())
	fmt.Println(testStack)
}

type Stack struct {
	stack []string
}

func (st *Stack) AddToStack(s string) {
	st.stack = append(st.stack, s)
}
func (st *Stack) PopFromStack() string {
	if len(st.stack) != 0 {
		returningString := st.stack[len(st.stack)-1]
		st.stack = st.stack[:len(st.stack)-1]
		return returningString
	}
	panic("The stack is empty")
}
