package goreload

import (
	"strings"
	"unicode"
)

func FixArticles(words []string) []string {

	//result := []string{}
	for i, wrd := range words {

		if wrd == "a" || wrd == "A" {

			first := []rune(words[i+1])[0]

			if strings.ContainsRune("aeiouhAEIOUH", first) {
				if unicode.IsUpper(first) {
					words[i] = strings.ToUpper(wrd) + "n"
				}else{
					words[i] = wrd + "n"
				}

			}
		}
	}

	return words
}
