package main

import "fmt"

func main() {
	// Maps devem ser inicializados
	// var approved map[int]string

	approved := make(map[int]string)

	approved[1234556461] = "Maria"
	approved[4534345135] = "João"

	fmt.Println(approved)

	for cpf, name := range approved {
		fmt.Printf("%s (CPF: %d)\n", name, cpf)
	}

	fmt.Println(approved[4534345135])

	delete(approved, 4534345135)

	fmt.Println(approved)
}
