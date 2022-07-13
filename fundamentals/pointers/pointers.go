package main

import "fmt"

func main() {
	// Go não possui aritmética de ponteiros
	i := 1

	var p *int = nil

	p = &i // pegando o endereço da variável

	*p++ // pegando o valor da variável
	// agora i = 2
	i++
	// agora i = 3

	fmt.Println(i, &i, p, *p, &p)
}
