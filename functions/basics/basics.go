package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Funções puras, determinísticas,não mexem com dados externos
	Funções impuras, não determinísticas, mexem com dados externos e podem ter vários e distintos resultados
*/

func f1() {
	fmt.Println("f1")
}

func f2() string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')

	return name
}

func f3(integer int) int {
	return integer
}

/*
	As funções em Go podem ter mais de um retorno
*/

func f4() (string, string) {
	return "Retorno 1", "Retorno 2"
}

func f5() (string, int, bool) {
	return "Retorno string", 4, true
}

func main() {
	fmt.Print(f2())
	fmt.Println(f3(7))
	fmt.Println(f4())
	fmt.Println(f5())
}
