package main

import (
	"fmt"
	"os"
)

func main() {
<<<<<<< HEAD

	name := ""
	if len(os.Args) == 1 {
		fmt.Println("Error")
	} else {
		for _, word := range os.Args[1:] {
			name = fmt.Sprintf("%v%v ", name, word)
		}
		fmt.Println(name + ", you have entered the Matrix. Welcome.")
	}
=======
	
	name := ""
	if len(os.Args) == 1 {
	  fmt.Println("Error")
	}else {
	  for _, word := range os.Args[1:]{
	    name = fmt.Sprintf("%v %v ", name, word)
	  }
	 fmt.Println(name + ", you have entered the Matrix. Welcome.")
  }
>>>>>>> 95826de17c2b4879da20ff3a9ae25fd6589b5e4e
}
