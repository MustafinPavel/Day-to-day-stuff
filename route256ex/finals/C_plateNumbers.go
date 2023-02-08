package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"regexp"
// 	"strconv"
// )

// //Разбирает строку на номерные знаки LDLL и LDDLL с помощью регулярки
// func plateNumbers() {
// 	//file1, _ := os.Open("input.txt")
// 	//r := bufio.NewReader(file1)
// 	r := bufio.NewReader(os.Stdin)
// 	out := bufio.NewWriter(os.Stdout)
// 	defer out.Flush()
// 	pat := regexp.MustCompile(`[A-Z]\d\d?[A-Z]{2}`)
// 	setsNumber := ReadIntFromLine(r)
// 	for l := 0; l < setsNumber; l++ {
// 		fullString := ReadOrdinaryString(r)
// 		numbers := pat.FindAllString(fullString, -1)
// 		fullString = pat.ReplaceAllString(fullString, "")
// 		if fullString == "" {
// 			for _, v := range numbers {
// 				fmt.Printf("%v ", v)
// 			}
// 			fmt.Printf("\n")
// 		} else {
// 			fmt.Println("-")
// 		}
// 	}

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
