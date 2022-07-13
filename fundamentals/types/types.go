package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	// uint8 uint16 uint32 uint64
	var a byte = 255

	fmt.Println("Tipo: ", reflect.TypeOf(a))

	// int8 int16 int32 int64
	b := math.MaxInt64
	fmt.Println("Tipo: ", reflect.TypeOf(b))

	// Representação inteira da tabela UNICODE
	var c rune = 'a'
	fmt.Println("Tipo: ", reflect.TypeOf(c))

	// float32 float64
	var d float32 = 49.99
	fmt.Println("Tipo: ", reflect.TypeOf(d))
	// padrão: float64
	fmt.Println("Tipo: ", reflect.TypeOf(49.99))

	// Booleans
	e := true
	fmt.Println("Tipo: ", reflect.TypeOf(e))

	// String
	f := "Hello, world!"
	fmt.Println("Tipo: ", reflect.TypeOf(f))
	g := `
		Múltiplas
		Linhas
		.
	`
	fmt.Println(g, len(g))

	// Char
	// Não há por padrão um valor char
	// O que se tem é uma forma de pegar o UNICODE
	// Para tanto, passa-se apenas uma única string entre aspas simples
	h := 'a'
	fmt.Println("Tipo: ", reflect.TypeOf(h))
}
