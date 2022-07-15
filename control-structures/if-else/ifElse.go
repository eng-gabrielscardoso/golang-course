package main

import "fmt"

func printResult(note float64) {
	if note >= 7 {
		fmt.Println("Approved; Note:", note)
	} else {
		fmt.Println("Not approved; Note:", note)
	}
}

func main() {
	printResult(5.6)
	printResult(7)
}
