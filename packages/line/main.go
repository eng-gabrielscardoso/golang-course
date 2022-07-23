package main

import "fmt"

// Run with: go run *.go
func main() {
	p1 := Point{2.0, 2.0}
	p2 := Point{2.0, 4.0}

	fmt.Println(cathetus(p1, p2))
	fmt.Println(Distance(p1, p2))
}
