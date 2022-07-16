package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3}
	s1 := []int{1, 2, 3}

	fmt.Printf("%v\n%v\n", a1, s1)
	fmt.Printf("%v\n%v\n", reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3]

	// Slice não é um array
	// Trata-se apenas de um ponteiro para o primeiro elemento dentro de outro array e o tamanho, para definir o seu range
	fmt.Printf("%v\n%v\n", a2, s2)
}
