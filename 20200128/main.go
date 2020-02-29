package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello world\n")

	// Two ways of defining variables
	// var name string = "michal"
	name := "marcello"

	if name == "michael" {
		fmt.Println("Hello michael")
	} else {
		// %s - put in the place of "%s" a string
		fmt.Printf("Hello beautiful %s\n", name)
	}
}
