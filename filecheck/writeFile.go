package filecheck

import (
	"fmt"
	"os"
	"strings"
)

func WriteOut(text string) string {

	output := os.Args[2]

	data := ReadFile()
	data = strings.ToUpper(string(data))

	//output, _ := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0644)

	err := os.WriteFile(output, []byte(data), 0644 )
	if err != nil {
		fmt.Println("error")
		return ""
	}
	return output
}
