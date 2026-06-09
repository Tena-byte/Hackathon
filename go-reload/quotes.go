package goreload

import "strings"

func FixQuotes(text string) string {

	wrds := strings.Split(text, "'")

	for i := 1; i < len(wrds); i+= 2{
		wrds[i] = strings.TrimSpace(wrds[i])
	}
	return strings.Join(wrds, "'")
}
