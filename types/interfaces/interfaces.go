package main

import "fmt"

type printable interface {
	toString() string
}

type person struct {
	name     string
	lastName string
}

type product struct {
	name  string
	price float64
}

func (p person) toString() string {
	return p.name + " " + p.lastName
}

func (p product) toString() string {
	return fmt.Sprintf("%s - U$ %.2f", p.name, p.price)
}

func print(x printable) {
	fmt.Println(x.toString())
}

func main() {
	var thing printable = person{"Gabriel", "Cardoso"}

	fmt.Println(thing.toString())

	thing = product{"Pencil", 1.96}

	fmt.Println(thing.toString())
}
