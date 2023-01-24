package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
)

func exJ() {
	//sTimer := time.Now()
	// file1, _ := os.Open("testFileInput")
	// r := bufio.NewReader(file1)
	wg := sync.WaitGroup{}
	r := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	dictionary := ScanLine(r)
	queries := ScanLine(r)
	var results []string = make([]string, len(queries)/2, len(queries))
	wg.Add(len(queries))
	for i := 0; i < len(queries); i++ {
		go func(i int) {
			testWord := queries[i]
			var longestSuffixWord string = ""
			var longestSuffix int
			for _, responseWord := range dictionary {
				for j := 1; j <= len(testWord); j++ {
					if j-longestSuffix > 1 {
						break
					}
					if longestSuffixWord == "" && testWord != responseWord {
						longestSuffixWord = responseWord
					}
					if len(testWord) >= j && len(responseWord) >= j && testWord != responseWord {
						if testWord[len(testWord)-j:] == responseWord[len(responseWord)-j:] && longestSuffix <= j {
							longestSuffixWord = responseWord
							longestSuffix = j
						}
					}
				}
			}
			results[i] = longestSuffixWord
			wg.Done()
		}(i)
	}
	wg.Wait()
	for _, v := range results {
		fmt.Fprintln(out, v)
	}
	//fmt.Printf("ExecTime: %v", time.Since(sTimer))
}

func ScanLine(r *bufio.Reader) []string {
	line, _, _ := r.ReadLine()
	dvStr := line
	dv, _ := strconv.Atoi(string(dvStr))
	dictionary := make([]string, 0, dv*2)
	for i := 0; i < dv; i++ {
		line, _, _ = r.ReadLine()
		nl := string(line)
		dictionary = append(dictionary, nl)
	}
	return dictionary
}
