package main

import "fmt"

func main() {
	s := make([]int, 10)

	fmt.Println(s)

	t := make([]int, 10, 20)

	// slice, tamanho do slice, capacidade do array interno referência do slice
	fmt.Println(t, len(t), cap(t))

	t = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	fmt.Println(t, len(t), cap(t))

	t = append(t, 1)

	// O slice é o array do JS kk
	fmt.Println(t, len(t), cap(t))
}
