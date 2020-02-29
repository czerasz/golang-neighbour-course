package main

import (
	"fmt"
)

func main() {
	greet("marcello")
	greet("michael")
	greet("marcello")
	greet("michael")
	greet("marcello")
	greet("michael")

	fmt.Printf("5+6 = %d\n", sum(5, 6))
}

func greet(name string) {
	if name == "michael" {
		fmt.Printf("Hello michael\n")
	} else {
		fmt.Printf("Hello beautify %s\n", name)
	}
}

func sum(a int, b int) int {
	return a + b
}
