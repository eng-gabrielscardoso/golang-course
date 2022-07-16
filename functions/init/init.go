package main

import (
	"fmt"
	"runtime/debug"
)

// A convensão em Go é de que há uma função Init que é chamada para fazer inicializações antes do main()

func init() {
	fmt.Println("Init")
}

func main() {
	fmt.Println("Main")
	debug.PrintStack()
}
