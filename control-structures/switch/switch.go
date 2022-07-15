package main

import (
	"fmt"
	"time"
)

func noteForConcept(n float64) string {
	var note = int(n)

	switch note {
	case 11:
		fallthrough
	case 10, 9:
		return "E"
	case 8, 7:
		return "B"
	case 6, 5:
		return "R"
	case 0, 1, 2, 3, 4:
		return "I"
	default:
		return "Invalid"
	}
}

func main() {
	fmt.Println(noteForConcept(10))
	fmt.Println(noteForConcept(8))
	fmt.Println(noteForConcept(5))
	fmt.Println(noteForConcept(4))

	t := time.Now()

	// Switch true
	switch {
	case t.Hour() < 12:
		fmt.Println("Bom dia")
	case t.Hour() < 18:
		fmt.Println("Boa tarde")
	default:
		fmt.Println("Boa noite")
	}
}
