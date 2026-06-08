package main

import (
	"fmt"
	"hackathon/filecheck"
	"hackathon/splitn"
	"os"
)

func main() {

	data := filecheck.ReadFile()
	result := filecheck.WriteOut(data)

	output, _ := os.ReadFile(result)

	fmt.Println(string(output))

	fmt.Println(splitn.Separate("Hello, world! How are you?"))

}
