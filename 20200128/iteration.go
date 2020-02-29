package main

import "fmt"

import "math"

func main() {
	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for _, day := range days {
		fmt.Printf("- %s\n", day)
	}

	fmt.Printf("--------------\n")

	for i, day := range days {
		if math.Mod(float64(i), 2.0) > 0.0 {
			fmt.Printf("- %d %s\n", i+1, day)
		} else {
			fmt.Printf("+ %d %s\n", i+1, day)
		}
	}
}
