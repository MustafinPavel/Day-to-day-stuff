package boringLecture

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

func main() {
	//ввод
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	favString := readLongLine(in)
	// логика
	result := make([][2]int, 130)

	for i := 0; i < len(favString); i++ {
		result[favString[i]][0] = int(favString[i])
		result[favString[i]][1] += (len(favString) - i) * (i + 1)
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})
	//вывод
	for i := 0; i < len(result); i++ {
		if result[i][0] > 0 {
			out.WriteString(string(result[i][0]) + ": " + strconv.Itoa(result[i][1]) + "\n")
		}
	}

}
func readLongLine(in *bufio.Reader) string {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 10000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	return string(tmpByteSlice)
}
