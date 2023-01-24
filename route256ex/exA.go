package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func exA() {
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n')
	s = strings.TrimSuffix(s, "\n")
	strCount, _ := strconv.Atoi(s)

	var input string
	for i := 0; i < int(strCount); i++ {
		in, _ := reader.ReadString('\n')
		input += in
	}
	sl := strings.Fields(input)
	for i := 0; i < len(sl); i++ {
		if i%2 == 0 {
			k, _ := strconv.Atoi(sl[i])
			k2, _ := strconv.Atoi((sl[i+1]))
			fmt.Println(k + k2)
		}
	}
}
