package main

import "fmt"

// Função pura
func inc1(n int) {
	n++
}

// Função impura
func inc2(n *int) {
	*n++
}

func main() {
	n := 1

	fmt.Println(n)

	inc1(n)

	fmt.Println(n)

	inc2(&n)

	fmt.Println(n)
}
