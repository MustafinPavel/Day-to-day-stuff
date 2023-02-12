package pullCards

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// choose optimal pull from a row of cards, taking one each time from one of the sides
func Pull() {
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	readOneInt(in) //skip line
	pulls := readOneInt(in)
	cards := readIntSlice(in)
	max := 0
	for i := 0; i < pulls+1; i++ {
		temp := 0
		for j := 0; j < i; j++ {
			temp += cards[j]
		}
		for j := 0; j < pulls-i; j++ {
			temp += cards[len(cards)-1-j]
		}
		if temp > max {
			max = temp
		}
	}
	out.WriteString(strconv.Itoa(max))
}
func readOneInt(in *bufio.Reader) int {
	temp, _, _ := in.ReadLine()
	tempInt, _ := strconv.Atoi(string(temp))
	return tempInt
}
func readIntSlice(in *bufio.Reader) []int {
	temp, _, _ := in.ReadLine()
	tempSl := strings.Split(string(temp), " ")
	var result []int = make([]int, 0, len(tempSl)*2)
	for _, v := range tempSl {
		num, _ := strconv.Atoi(v)
		result = append(result, num)
	}
	return result
}
