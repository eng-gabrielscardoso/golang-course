package main

import "fmt"

type note float64

func (n note) between(begin, end float64) bool {
	return float64(n) >= begin && float64(n) <= end
}

func noteForConcept(n note) string {
	if n.between(9.0, 10.0) {
		return "E"
	} else if n.between(8.0, 8.99) {
		return "B"
	} else if n.between(7.0, 7.99) {
		return "R"
	} else {
		return "I"
	}
}

func main() {
	fmt.Println(noteForConcept(9.8))
	fmt.Println(noteForConcept(8))
	fmt.Println(noteForConcept(7))
	fmt.Println(noteForConcept(5))
}
