package exC

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// ужасное решение
func Solution() {
	scanner := bufio.NewScanner(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	lines := make([]string, 0, 1000)
	scanner.Scan()
	pairNumbeString := scanner.Text()
	pairNumber, _ := strconv.Atoi(pairNumbeString)
	for i := 0; i < pairNumber*2; i++ {
		scanner.Scan()
		newLine := scanner.Text()
		lines = append(lines, newLine)
	}
	for i, v := range lines {
		if i%2 == 1 {
			proggersNumber, _ := strconv.Atoi(lines[i-1])
			var skillList = make([]int, 0, proggersNumber)
			for _, v := range strings.Fields(v) {
				skill, _ := strconv.Atoi(v)
				skillList = append(skillList, skill)
			}
			for t := 0; t < len(skillList)/2; t++ {
				var firstProgIndex int = -1
				var firstProgValue int = -1
				var bestPairDif int = -1
				var bestPairIndex int = -1
				for i := 0; i < len(skillList); i++ {
					if firstProgIndex == -1 && skillList[i] != 0 {
						firstProgIndex = i
						firstProgValue = skillList[i]
						skillList[i] = 0
						break
					}
				}
				for k, v := range skillList {
					if v != 0 {
						dif := int(math.Abs(float64(firstProgValue - v)))
						if bestPairDif == -1 || dif < bestPairDif {
							bestPairDif = dif
							bestPairIndex = k
						}
					}
					if k == len(skillList)-1 {
						fmt.Fprintln(out, firstProgIndex+1, bestPairIndex+1)
						skillList[bestPairIndex] = 0
					}
				}
			}
			fmt.Fprintln(out, "")
		}
	}
}
