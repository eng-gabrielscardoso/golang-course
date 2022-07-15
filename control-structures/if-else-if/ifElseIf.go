package main

import (
	"fmt"
)

func noteForConcept(note float64) string {
	if note >= 9 && note <= 10 {
		return "E"
	} else if note >= 8 && note < 9 {
		return "B"
	} else if note >= 5 && note < 8 {
		return "R"
	} else if note >= 0 && note < 5 {
		return "I"
	} else {
		return "Invalid note"
	}
}

func printResult() {
	var result string = noteForConcept(10)

	var message string

	if result == "E" {
		message = "Congratulations! Concept:" + result
	} else if result == "B" || result == "R" {
		message = "Good! Concept:" + result
	} else if result == "I" {
		message = "You must study more! Concept:" + result
	} else {
		message = "Error! Invalid concept has been taken"
	}

	fmt.Println(message)
}

func main() {
	printResult()
}
