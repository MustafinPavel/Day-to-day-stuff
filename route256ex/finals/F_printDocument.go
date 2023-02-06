package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"sort"
// 	"strconv"
// 	"strings"
// )

// //Перечисляет, какие страницы документа необходимо допечатать и выдаёт в корректном формате.
// func printDocument() {
// 	//file1, _ := os.Open("input.txt")
// 	//r := bufio.NewReader(file1)
// 	r := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()
// 	globalSetsAmount := ReadIntFromLine(r)
// 	for i := 0; i < globalSetsAmount; i++ {
// 		pagesOfTheDocument := ReadIntFromLine(r)
// 		tasksString := ReadOrdinaryString(r)
// 		printedPages := make(map[int]int)
// 		separatedTasks := strings.Split(tasksString, ",")
// 		for _, v := range separatedTasks {
// 			if strings.Contains(v, "-") {
// 				minMax := strings.Split(v, "-")
// 				min, _ := strconv.Atoi(minMax[0])
// 				max, _ := strconv.Atoi(minMax[1])
// 				for t := min; t <= max; t++ {
// 					printedPages[t] = 1
// 				}
// 			} else {
// 				t, _ := strconv.Atoi(v)
// 				printedPages[t] = 1
// 			}
// 		}
// 		corSl := CreateCorrectSlice(printedPages, pagesOfTheDocument)
// 		corStr := CreateShortString(corSl)
// 		fmt.Println(corStr)
// 		//out.WriteString(corStr)
// 	}
// }

// func CreateShortString(sl []int) string {
// 	res := ""
// 	temp := ""
// 	var isPeriod bool = false
// 	for i := 0; i < len(sl); i++ {
// 		if i == 0 {
// 			temp += strconv.Itoa(sl[i])
// 			continue
// 		}
// 		if sl[i]-sl[i-1] == 1 {
// 			isPeriod = true
// 		}
// 		if sl[i]-sl[i-1] > 1 && isPeriod {
// 			temp += "-"
// 			temp += strconv.Itoa(sl[i-1])
// 			res += temp
// 			res += ","
// 			temp = ""
// 			temp += strconv.Itoa(sl[i])
// 			isPeriod = false
// 			continue
// 		}
// 		if sl[i]-sl[i-1] > 1 && !isPeriod {
// 			res += temp
// 			res += ","
// 			temp = ""
// 			temp += strconv.Itoa(sl[i])
// 		}
// 		if isPeriod && i == len(sl)-1 {
// 			temp += "-"
// 			temp += strconv.Itoa(sl[i])
// 			res += temp
// 			temp = ""
// 			isPeriod = false
// 		}
// 	}
// 	res += temp
// 	return res
// }
// func CreateCorrectSlice(m map[int]int, maxPages int) []int {
// 	var tempSlice sort.IntSlice
// 	for k, _ := range m {
// 		tempSlice = append(tempSlice, k)
// 	}

// 	sort.Sort(tempSlice)
// 	resultingSlice := make([]int, 0, len(tempSlice))
// 	for i := 1; i <= maxPages; i++ {
// 		if !sliceContains(tempSlice, i) {
// 			resultingSlice = append(resultingSlice, i)
// 		}
// 	}
// 	return resultingSlice
// }
// func ReadOrdinaryString(r *bufio.Reader) string {
// 	line, _, _ := r.ReadLine()
// 	return string(line)
// }
// func ReadIntFromLine(r *bufio.Reader) int {
// 	line, _, _ := r.ReadLine()
// 	lineInt, _ := strconv.Atoi(string(line))
// 	return lineInt
// }
// func sliceContains(sl []int, i int) bool {
// 	for _, v := range sl {
// 		if v == i {
// 			return true
// 		}
// 	}
// 	return false
// }
