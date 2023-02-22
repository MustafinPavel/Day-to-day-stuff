package pupilsAndTest

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	pups := readSingleInt(in)
	vars := readSingleInt(in)
	deck := readSingleInt(in)
	petyasSide := readSingleInt(in)
	//logic
	petyasPos := 2*(deck-1) + petyasSide
	var vasyasSide, vasyasPos int
	if vars%2 == 0 {
		vasyasSide = petyasSide
		switch {
		case petyasPos+vars <= pups:
			vasyasPos = petyasPos + vars
		case petyasPos-vars >= 1:
			vasyasPos = petyasPos - vars
		default:
			vasyasPos = -1
		}
	}
	if vars%2 == 1 {
		vasyasSide = notPetyasSide(petyasSide)
		switch {
		case petyasSide == 2 && petyasPos-vars >= 1:
			vasyasPos = petyasPos - vars
		case petyasSide == 2 && petyasPos+vars <= pups:
			vasyasPos = petyasPos + vars
		case petyasSide == 1 && petyasPos+vars <= pups:
			vasyasPos = petyasPos + vars
		case petyasSide == 1 && petyasPos-vars >= 1:
			vasyasPos = petyasPos - vars
		default:
			vasyasPos = -1
		}
	}
	//output
	vasyasDeck := (vasyasPos-vasyasSide)/2 + 1
	if vasyasDeck > 0 {
		out.WriteString(strconv.Itoa(vasyasDeck) + " " + strconv.Itoa(vasyasSide))
	} else {
		out.WriteString(strconv.Itoa(vasyasPos))
	}
}
func notPetyasSide(petyasSide int) int {
	if petyasSide == 2 {
		return 1
	} else {
		return 2
	}
}

func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}

//Если вариантов чётное число, то в приоритете надо садиться сзади Пети на тот же ряд
// (при условии, что хватит учеников), в противном случае спереди на тот же ряд.
//Если вариантов нечётное число, то в приоритете надо садиться спереди на другой ряд,
// (при условии, что хватит учеников), в противном случае сзади на другой ряд.
//Если при нечётном варианте Петя сидит на [1], то Вася в приоритете сзади на [2]
//Если при нечётном варианте Петя сидит на [2], то Вася в приоритете спереди на [1]
