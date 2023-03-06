package operationSystemsLite

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
	readSingleInt(in) //skip line
	osAmount := readSingleInt(in)
	OSs := make([][2]int, 0, osAmount)
	for i := 0; i < osAmount; i++ {
		OSs = append(OSs, readTwoInts(in))
	}
	//логика
	workingOSs := make(map[[2]int]bool)
	var counter int
	for _, installing := range OSs {
		for installed, _ := range workingOSs {
			if installing[1] >= installed[0] && installing[0] <= installed[1] {
				counter--
				delete(workingOSs, installed)
			}
		}
		workingOSs[installing] = true
		counter++
	}
	//вывод
	out.WriteString(strconv.Itoa(counter))

}
func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}
func readTwoInts(r *bufio.Reader) (res [2]int) {
	line, _, _ := r.ReadLine()
	str := strings.Split(string(line), " ")

	for i := 0; i < len(str); i++ {
		num, _ := strconv.Atoi(str[i])
		res[i] = num
	}
	return
}
