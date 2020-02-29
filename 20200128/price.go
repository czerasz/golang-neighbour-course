package main

import "fmt"

func main() {
	product := "t-shirt"
	color := "red"
	var price int

	if product == "t-shirt" && color == "blue" {
		price = 40
	} else if product == "t-shirt" && color == "red" {
		price = 100
	} else {
		price = 10
	}

	if price > 80 {
		fmt.Printf("Hey You deserve a discount!\n")
		price = price - 10
	}

	if product == "t-shirt" || product == "sock" {
		price = price - 1
	}

	fmt.Printf("Please pay %d EUR\n", price)
}
