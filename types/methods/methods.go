package main

import (
	"fmt"
	"strings"
)

type person struct {
	name     string
	lastName string
}

func (p person) getFullName() string {
	return p.name + " " + p.lastName
}

func (p *person) setFullName(fullName string) {
	parts := strings.Split(fullName, " ")

	p.name = parts[0]
	p.lastName = parts[1]
}

func main() {
	// Também poderia ser inicializado sem valores: person{}
	person1 := person{
		name:     "Gabriel",
		lastName: "Cardoso",
	}

	fmt.Println(person1.getFullName())

	person1.setFullName("Fabiana Barreto")

	fmt.Println(person1.getFullName())
}
