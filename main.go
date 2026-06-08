package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2{
		fmt.Println("Uasage : go run . <filename>")
		return
	}

	file := os.Args[1]
	files, _ := os.ReadFile(file)


	if _, err := os.Stat(file); err != nil{
		fmt.Println("file not found")
		return
	}

	fmt.Println(string(files))

	


}
