package abracadabra

import (
	"bufio"
	"fmt"
	"os"
)

// Solution is in Another Castle.
func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	fmt.Println(in) //plug
}
