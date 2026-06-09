package main

import (
	"fmt"
	goreload "hackathon/go-reload"
)

func main() {

	// data := filecheck.ReadFile()
	// result := filecheck.WriteOut(data)

	// output, _ := os.ReadFile(result)

	// fmt.Println(string(output))

	//fmt.Println(splitn.Separate("Hello, world! How are you?"))

	fmt.Println(goreload.CapitalizeFirstChar("hello, world!"))

	fmt.Println(goreload.CapitalizeN([]string{"hello,", "world!", "Terna"}, 1))

}
