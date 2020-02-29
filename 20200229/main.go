package main

import (
	"fmt"
)

type Person struct {
	name   string
	gender string
	age    int
}

func (person Person) printAge() {
	fmt.Printf("%s's age is: %d\n", person.name, person.age)
}

func (p Person) printGender(c string) {
	if c == "capital" {
		fmt.Printf("THE GENDER IS: %s\n", p.gender)
	} else {
		fmt.Printf("the gender is: %s\n", p.gender)
	}
}

func (p Person) sing(c string) {
	fmt.Printf("lalalala\n")
}

type Animal struct {
	name   string
	gender string
	age    int
}

// This method doesn't belong to animals
// func (a Animal) sing(f string) {
// 	fmt.Printf("loar\n")
// }

type Human interface {
	sing(c string)
}

func humanTest(h Human) {
	h.sing("high")
}

func main() {
	person1 := Person{
		name:   "Karl",
		gender: "male",
		age:    23,
	}

	person1.printAge()

	person2 := Person{
		name:   "Joe",
		gender: "male",
		age:    44,
	}
	person2.printAge()
	person2.printGender("asd")

	animal := Animal{
		name:   "Rex",
		gender: "male",
		age:    4,
	}

	fmt.Printf("Rex's age is: %d\n", animal.age)

	humanTest(person1)

	// This line will fail
	humanTest(animal)
}
