package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Person struct {
	ID     int
	Name   string
	Age    int
	Height float32
}

func main() {
	file, err := os.Open("read.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close() // we will learn what defer is at some later point

	count := 0
	ageSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ";")

		id, _ := strconv.Atoi(fields[0])
		age, _ := strconv.Atoi(fields[2])

		human := Person{
			ID:   id,
			Name: fields[1],
			Age:  age,
		}

		fmt.Printf("ID: %d, Name: '%s', Age: %d, Height: %.2f\n", human.ID+5, human.Name, human.Age-10, human.Height)
		count = count + 1
		ageSum = ageSum + human.Age
	}

	fmt.Printf("Average age: %f\n", (float32(ageSum) / float32(count)))

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
