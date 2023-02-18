package lineSimplification

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	inputLine := readLongLine(in)

	counter := 1
	for i := 0; i < len(inputLine); i++ {
		if i > 0 {
			cur := string(inputLine[i])
			prev := string(inputLine[i-1])
			if prev == cur {
				counter++
			} else {
				addToResult(prev, counter, out)
				counter = 1
			}
			if i == len(inputLine)-1 {
				addToResult(cur, counter, out)
			}
		}
	}
}
func addToResult(a string, i int, out *bufio.Writer) {
	if i > 1 {
		out.WriteString(a + strconv.Itoa(i))
	} else {
		out.WriteString(a)
	}
}

func readLongLine(in *bufio.Reader) string {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	return string(tmpByteSlice)
}
