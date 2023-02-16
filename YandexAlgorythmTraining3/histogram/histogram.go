package histogram

import (
	"bufio"
	"os"
	"sort"
)

// Рисует гистограмму частоты встречаемых символов в загружаемом тексте
func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var tmpByteSlice []byte
	tmpByteSlice, _ = in.ReadBytes('_')
	//logic
	histoDataMap := make(map[string]int)
	fillDataMap(tmpByteSlice, histoDataMap)
	histoHeight := defineHistoHeight(histoDataMap)
	histoWidth := len(histoDataMap)
	drawHistogram(histoDataMap, histoHeight, histoWidth, out)
}

func fillDataMap(bs []byte, dataMap map[string]int) {
	for _, v := range string(bs) {
		if v != '\n' && v != ' ' {
			dataMap[string(v)] += 1
		}
	}
}

func defineHistoHeight(dataMap map[string]int) (res int) {
	for _, v := range dataMap {
		if v > res {
			res = v
		}
	}
	res += 1
	return
}

func drawHistogram(data map[string]int, height, width int, out *bufio.Writer) {
	rightOrderSlice := make([]string, 0, len(data))
	for k := range data {
		rightOrderSlice = append(rightOrderSlice, k)
	}
	sort.Strings(rightOrderSlice)
	for h := 1; h <= height; h++ {
		for w := 0; w < width; w++ {
			currentSymbol := rightOrderSlice[w]
			if h == height {
				out.WriteString(currentSymbol)
				continue
			}
			if data[currentSymbol] >= height-h {
				out.WriteString("#")
			} else {
				out.WriteString(" ")
			}
		}
		out.WriteString("\n")
	}
}
