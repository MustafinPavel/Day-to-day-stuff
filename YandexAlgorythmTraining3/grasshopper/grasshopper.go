package grasshopper

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	//ввод
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	plank, hop := readTwoInts(in)
	//логика
	resultsSlice := countVariations(plank, hop)
	//вывод
	finalResult := resultsSlice[len(resultsSlice)-1]
	out.WriteString(strconv.Itoa(finalResult))
}

func countVariations(plank, hop int) []int {

	var variations []int = make([]int, plank)
	if plank == 1 {
		return []int{1}
	} else {
		variations[0], variations[1] = 1, 1
	}
	for i := 2; i < plank; i++ {
		input := 0
		for j := 0; j < hop; j++ {
			if i-j-1 >= 0 {
				input += variations[i-1-j]
			} else {
				input += 0
			}
		}
		variations[i] = input
	}
	return variations
}

func readTwoInts(r *bufio.Reader) (plank, hop int) {
	result := make([]int, 0, 10)
	line, _, _ := r.ReadLine()
	slice := strings.Fields(string(line))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result[0], result[1]
}
