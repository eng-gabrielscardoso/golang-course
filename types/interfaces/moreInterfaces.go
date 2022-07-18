package main

import "fmt"

type turbo interface {
	activeTurbo()
}

type ferrari struct {
	model          string
	activedTurbo   bool
	actualVelocity int
}

func (f *ferrari) activeTurbo() {
	f.activedTurbo = true
}

func main() {
	f := ferrari{"F30", false, 0}

	f.activeTurbo()

	var f2 turbo = &ferrari{"F40", false, 0}

	f2.activeTurbo()

	fmt.Println(f, f2)
}
