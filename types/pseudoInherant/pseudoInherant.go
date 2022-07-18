package main

import "fmt"

type car struct {
	name           string
	actualVelocity int
}

type ferrari struct {
	// campos anonimos
	car
	onTurb bool
}

func main() {
	f := ferrari{}
	f.name = "Ferrari F40"

	fmt.Println(f.name)

	f.onTurb = true

	fmt.Println(f)

}
