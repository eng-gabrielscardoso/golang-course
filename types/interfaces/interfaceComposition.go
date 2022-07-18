package main

import "fmt"

type sport interface {
	activeTurbo()
}

type suit interface {
	doBalize()
}

type sportSuit interface {
	sport
	suit
}

type bmw struct{}

func (b bmw) activeTurbo() {
	fmt.Println("Turbo active")
}

func (b bmw) doBalize() {
	fmt.Println("Doing balize")
}

func main() {
	b := bmw{}

	b.activeTurbo()
	b.doBalize()
}
