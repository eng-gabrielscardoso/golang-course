package main

import (
	"fmt"
	"math"
)

func main() {
	// Forma padrão
	const PI float64 = 3.14159265

	var raio = 3.2

	// Forma reduzida
	area := PI * math.Pow(raio, 2)

	// Obrigatoriamente, toda variável em Go deve ser usada
	fmt.Println(area)

	const (
		A int64 = 1
		B int64 = 2
		C int64 = 3
	)

	var (
		D int64 = 1
		E int64 = 2
		F int64 = 3
	)

	fmt.Println(D, E, F)

	var g, h bool = true, false

	fmt.Println(g, h)

	i, j := "Uepa!", false

	fmt.Println(i, j)
}
