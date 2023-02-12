package bracketEncode

import "strings"

// Encodes a word with a bracket sequence. If the letter occurs once it's "(" and
// if the letter occurs twice or more times it's ")"
func DuplicateEncode(word string) string {
	temp := make(map[rune]int)
	word = strings.ToLower(word)
	var result string
	for _, v := range word {
		if _, ok := temp[v]; ok {
			temp[v] += 1
		} else {
			temp[v] = 1
		}
	}
	for _, v := range word {
		if temp[v] == 1 {
			result += "("
		} else {
			result += ")"
		}
	}
	return result
}
