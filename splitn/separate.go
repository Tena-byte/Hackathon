package splitn

import (
	"strings"
	"unicode"
)

func Separate(text string) []string {

	words := strings.Fields(text)
	result := []string{}

	for _, r := range words {

		for _, ch := range r {

			if unicode.IsPunct(ch) {
				s := string(ch)
				wrd := strings.NewReplacer(s, " "+s).Replace(r)
				result = append(result, wrd)
			}
		}

	}

	return result

}
