package main

import "fmt"

type course struct {
	name string
}

func main() {
	var thing interface{}

	fmt.Println(thing)

	thing = 3

	fmt.Println(thing)

	type dynamic interface{}

	var thing2 dynamic = "uepa"

	fmt.Println(thing2)

	thing2 = 3

	fmt.Println(thing2)

	// Nesse caso, a interface poderia ser utilizada como tipo genérico
}
