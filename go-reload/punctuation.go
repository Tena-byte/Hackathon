package goreload

import (
	"strings"
	"unicode"
)

func FixPuntuation(text string) string {

	for _, r := range text {

		if unicode.IsPunct(r) {
			s := string(r)
			text = strings.NewReplacer(" "+s, s, ""+s, s).Replace(text)
		}
	}
	return text
}
