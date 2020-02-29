package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Hello sir\n")
	fmt.Printf("cos(0) = %.3f\n", math.Cos(0))

	fmt.Printf("|-1| = %.3f\n", math.Abs(-1))
	fmt.Printf("|-5| = %.3f\n", math.Abs(-5))
	fmt.Printf("|5| = %.3f\n", math.Abs(5))

	fmt.Printf("%.3f\n", math.Ceil(1.1))
	fmt.Printf("min: %.3f\n", math.Min(20, 9))
}
