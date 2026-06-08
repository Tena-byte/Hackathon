package main

import (
	"fmt"
	"hackathon/filecheck"
	"os"
)

func main() {

	
	data := filecheck.ReadFile()
	result := filecheck.WriteOut(data)
	
	output, _ := os.ReadFile(result)

	fmt.Println(string(output))
}
