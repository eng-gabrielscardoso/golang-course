package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Divisão:", a/b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Resto: ", a%b)
	fmt.Println("Conjunção Bit a bit: ", a&b)
	fmt.Println("Disjunção Bit a bit: ", a|b)
	fmt.Println("Disjunção Exclusiva Bit a bit: ", a^b)

	c := 3.0
	d := 2.0

	fmt.Println("Maior:", math.Max(float64(c), float64(d)))
	fmt.Println("Menor:", math.Min(float64(c), float64(d)))
	fmt.Println("Potenciação:", math.Pow(float64(c), float64(d)))
}
