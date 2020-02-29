package main

import (
	"fmt"
)

type Person struct {
	ID int
	Name string
	Age int
	Height float32
}

func main() {
	var name string = "Arnold 123"
	var arnold Person = Person {
		ID: 5,
		Name: name,
		Age: 32+9,
	}

	var beti Person = Person {
		ID: 2,
		Name: "Beti",
		Age: 3+9,
		Height: 182.1234,
	}

	fmt.Println(arnold.ID)
	fmt.Println(arnold.Name)
	fmt.Println(arnold.Age)
	fmt.Println(arnold.Height)

	fmt.Println("---------------------")

	fmt.Println(beti.ID)
	fmt.Println(beti.Name)
	fmt.Println(beti.Age)
	fmt.Println(beti.Height)

	fmt.Println("---------------------")

	fmt.Printf("ID: %d, Name: '%s', Age: %d, Height: %.2f\n", beti.ID, beti.Name, beti.Age, beti.Height)
}
