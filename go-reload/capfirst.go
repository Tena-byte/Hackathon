package goreload

import (
	"strings"
)


// Write a function capitalize(s string) string that capitalizes only the first letter of a word 
// (making the rest lowercase). Then write capitalizeN(words []string, n int) []string 
// that capitalizes the last N words in the slice. This mimics the (cap, N) modifier in go-reloaded.




func CapitalizeFirstChar(text string) string {
	return strings.ToUpper(string(text[0]))
}

func CapitalizeN(words []string, n int) []string {
	if n <= 0 {
		return words
	}

	//result := []string{}
	l := len(words)

	if n > l {
		n = l
	}

	start := l - n

	for i := start; i < l; i++ {
		words[i] = strings.ToUpper(words[i])
	}

	return words
}
