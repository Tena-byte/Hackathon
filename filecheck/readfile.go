package filecheck

import (
	"fmt"
	"os"
)

func ReadFile() string{

	if len(os.Args) != 3 {
		fmt.Println("Uasage : go run . <filename> output.txt" )
		return ""
	}

	file := os.Args[1]
	files, _ := os.ReadFile(file)

	if _, err := os.Stat(file); err != nil {
		fmt.Println("file not found")
		return ""
	}

	return string(files)
}
