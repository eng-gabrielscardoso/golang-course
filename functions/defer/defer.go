package main

import "fmt"

func getValue(n int) int {
	defer fmt.Println("Defer")

	if n >= 5000 {
		fmt.Println("Major")
		return 5000
	} else {
		fmt.Println("Minor")
		return n
	}
}

func main() {
	fmt.Println(getValue(6000))
	fmt.Println()
	fmt.Println(getValue(6))
}
