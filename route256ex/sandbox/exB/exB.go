package exB

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Ужасное решение
func Solution() {
	// file1, _ := os.Open("testFileInput")
	// r := bufio.NewReader(file1)
	r := bufio.NewReader(os.Stdin)
	result := bufio.NewWriter(os.Stdout)
	defer result.Flush()
	lines := make([]string, 0, 20)
	a, _, _ := r.ReadLine()
	purchaseNumbeString := string(a)
	purchaseNumber, _ := strconv.Atoi(purchaseNumbeString)
	for i := 0; i < purchaseNumber*2; i++ {
		a, _, _ = r.ReadLine()
		if len(string(a)) != 0 {
			lines = append(lines, string(a))
		}
	}
	for i, v := range lines {
		if i%2 == 1 {
			var goodsSlice []string = make([]string, 0, 200000)
			goodsSlice = strings.Fields(v)
			fmt.Printf("goods amount: %v\n", len(goodsSlice))
			keysMap := map[int]int{}
			for _, v := range goodsSlice {
				good, _ := strconv.Atoi(v)
				if _, ok := keysMap[good]; ok {
					keysMap[good] += 1
				} else {
					keysMap[good] = 1
				}
			}
			var outputSum int
			for k, v := range keysMap {
				outputSum += k * (v - v/3)
			}
			fmt.Fprintln(result, outputSum)
			fmt.Printf("right answer is: %v\n", 9531*(200000-200000/3))
		}
	}
}
