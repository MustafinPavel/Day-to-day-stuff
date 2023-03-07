package minimalRectangle

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
	dotsAmount := readSingleInt(in)
	//логика
	var (
		minX int
		minY int
		maxX int
		maxY int
	)
	for i := 0; i < dotsAmount; i++ {
		x, y := readPointCoords(in)
		if i == 0 {
			minX, maxX, minY, maxY = x, x, y, y
		}
		if x < minX {
			minX = x
		}
		if x > maxX {
			maxX = x
		}
		if y < minY {
			minY = y
		}
		if y > maxY {
			maxY = y
		}
	}
	//вывод
	out.WriteString(strconv.Itoa(minX) + " " + strconv.Itoa(minY) + " " + strconv.Itoa(maxX) + " " + strconv.Itoa(maxY))

}
func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
func readPointCoords(r *bufio.Reader) (x, y int) {
	line, _, _ := r.ReadLine()
	str := strings.Split(string(line), " ")
	x, _ = strconv.Atoi(str[0])
	y, _ = strconv.Atoi(str[1])
	return
}
