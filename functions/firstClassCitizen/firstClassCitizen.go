package main

import "fmt"

var sum = func(a, b int) int { return a + b }

func main() {
	fmt.Println(sum(2, 3))
}
