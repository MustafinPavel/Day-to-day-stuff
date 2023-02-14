package cardCounter

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// Choose optimal pull set from the row of cards, pulling one card from one of the sides at a time
func main() {
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	cardsAmount := ReadOneInt(in)
	pulls := ReadOneInt(in)
	if pulls > cardsAmount {
		pulls = cardsAmount
	}
	//считываем []byte
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
	for isNotEnded {
		cds, end := ReadFullLine(in)
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, cds...)
	}
	//превращаем в []int
	tmpStrSlice := strings.Split(string(tmpByteSlice), " ")
	cards := make([]int, 0, len(tmpStrSlice)*2)
	for _, v := range tmpStrSlice {
		num, _ := strconv.Atoi(v)
		cards = append(cards, num)
	}
	//алгоритм выбора комбинации
	max := 0
	for i := 0; i < pulls; i++ {
		max += cards[i]
	}
	temp := max
	for i := 0; i < pulls; i++ {
		temp += cards[len(cards)-1-i]
		temp -= cards[pulls-1-i]
		if temp > max {
			max = temp
		}
	}
	//вывод
	out.WriteString(strconv.Itoa(max))
}

// функции считывания входящих данных
func ReadOneInt(in *bufio.Reader) int {
	temp, _, _ := in.ReadLine()
	tempInt, _ := strconv.Atoi(string(temp))
	return tempInt
}
func ReadFullLine(in *bufio.Reader) ([]byte, bool) {
	temp, isNotEnded, _ := in.ReadLine()
	return temp, isNotEnded
}
