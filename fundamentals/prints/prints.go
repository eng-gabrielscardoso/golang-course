package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.")
	fmt.Print("\n")

	fmt.Println("Com quebra de linha.")

	x := 3.141516

	// fmt.Println("O valor é " + x)

	xString := fmt.Sprint(x)

	fmt.Println("O valor é " + xString)
	fmt.Println("O valor é", x)
	fmt.Printf("O valor é %.2f.\n", x)
}
