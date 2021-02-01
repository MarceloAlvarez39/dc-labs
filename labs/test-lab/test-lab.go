package main

import (
	"fmt"
	"os"
)

func main() {

	name := ""
	if len(os.Args) == 1 {
		fmt.Println("Error")
	} else {
		for _, word := range os.Args[1:] {
			name = fmt.Sprintf("%v%v ", name, word)
		}
		fmt.Println(name + ", you have entered the Matrix. Welcome.")
	}
}
