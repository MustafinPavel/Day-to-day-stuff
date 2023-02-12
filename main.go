package main

func main() {

}

// func MyReverse(s []string) {
// 	for i := 0; i < len(s)/2; i++ {
// 		first := s[i]
// 		last := s[len(s)-1-i]
// 		s[len(s)-1-i] = first
// 		s[i] = last
// 	}
// }

// func InsertionSort(array []int) {
// 	for i := 0; i < len(array); i++ {
// 		j := i
// 		temp := array[i]
// 		for j > 0 && array[j-1] > temp {
// 			array[j] = array[j-1]
// 			j--
// 		}
// 		array[j] = temp
// 	}
// }

// type Stack struct {
// 	stack []string
// }
// func (st *Stack) AddToStack(s string) {
// 	st.stack = append(st.stack, s)
// }
// func (st *Stack) PopFromStack() string {
// 	if len(st.stack) != 0 {
// 		returningString := st.stack[len(st.stack)-1]
// 		st.stack = st.stack[:len(st.stack)-1]
// 		return returningString
// 	}
// 	panic("The stack is empty")
// }
