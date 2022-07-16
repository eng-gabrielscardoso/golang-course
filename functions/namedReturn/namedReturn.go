package main

import "fmt"

func swap(p1, p2 int) (second int, first int) {
	// Quando os valores dos retornos são definidos, não é necessário passá-los no return
	// NÃO É EARLY RETURN
	second = p2
	first = p1

	// Clean return
	return
}

func main() {
	fmt.Println(swap(1, 2))
}
