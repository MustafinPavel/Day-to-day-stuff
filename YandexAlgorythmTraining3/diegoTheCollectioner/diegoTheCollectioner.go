package diegoTheCollectioner

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	//input
	file1, _ := os.Open("input.txt")
	in := bufio.NewReader(file1)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	readSingleInt(in) //skip DiegoCardsAmount line
	cards := readLongIntSlice(in)
	readSingleInt(in) //skip guestsAmount line
	guestsLimits := readLongIntSlice(in)
	//logic
	cardsSorted := createSortedUniqueCardsSlice(cards)
	guestsSorted := createSortedGuestsSlice(guestsLimits) //[0]index, [1]cardLimit, [2]answer
	cc := 0                                               //card counter
	currentGuestIdx := 0
	for i := 0; i < len(cardsSorted); i++ {
		currentCard := cardsSorted[i]
		for len(guestsSorted) > currentGuestIdx && currentCard >= guestsSorted[currentGuestIdx][1] {
			guestsSorted[currentGuestIdx][2] = cc
			currentGuestIdx++
		}
		cc++
		if i == len(cardsSorted)-1 {
			for currentGuestIdx < len(guestsSorted) {
				guestsSorted[currentGuestIdx][2] = cc
				currentGuestIdx++
			}
		}
	}
	sort.Slice(guestsSorted, func(i, j int) bool {
		return guestsSorted[i][0] < guestsSorted[j][0]
	})
	//output
	for _, res := range guestsSorted {
		out.WriteString(strconv.Itoa(res[2]) + "\n")
	}
}

func createSortedGuestsSlice(guestsLimits []int) [][]int {
	guestsSorted := make([][]int, 0, len(guestsLimits))
	for i := 0; i < len(guestsLimits); i++ {
		guestsSorted = append(guestsSorted, []int{i, guestsLimits[i], 0})
	}
	sort.Slice(guestsSorted, func(i, j int) bool {
		return guestsSorted[i][1] < guestsSorted[j][1]
	})
	return guestsSorted
}

func createSortedUniqueCardsSlice(cards []int) []int {
	cardsMap := make(map[int]int)
	cardsSorted := make([]int, 0, len(cards))
	for i := 0; i < len(cards); i++ {
		cardsMap[cards[i]]++
		if cardsMap[cards[i]] == 1 {
			cardsSorted = append(cardsSorted, cards[i])
		}
	}
	sort.Ints(cardsSorted)
	return cardsSorted
}

func readSingleInt(r *bufio.Reader) int {
	line, _, _ := r.ReadLine()
	lineInt, _ := strconv.Atoi(string(line))
	return lineInt
}

func readLongIntSlice(in *bufio.Reader) []int {
	isNotEnded := true
	tmpByteSlice := make([]byte, 0, 100000)
	result := make([]int, 0, 100000)
	for isNotEnded {
		tmp, end, _ := in.ReadLine()
		isNotEnded = end
		tmpByteSlice = append(tmpByteSlice, tmp...)
	}
	slice := strings.Fields(string(tmpByteSlice))
	for i := 0; i < len(slice); i++ {
		t, _ := strconv.Atoi(slice[i])
		result = append(result, t)
	}
	return result
}
