package filecheck

import (
	"fmt"
	"os"
	"strings"
)

func WriteOut(text string) string {

	output := os.Args[2]
	data := strings.ToUpper(ReadFile())

	err := os.WriteFile(output, []byte(data), 0644)
	if err != nil {
		fmt.Println("error")
		return ""
	}
	return output
}
