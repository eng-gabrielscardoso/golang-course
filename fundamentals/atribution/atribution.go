package main

import "fmt"

func main() {
	var a byte = 3
	fmt.Println(a)

	b := 3
	fmt.Println(b)
	b += 3
	fmt.Println(b)
	b -= 3
	fmt.Println(b)
	b /= 3
	fmt.Println(b)
	b *= 3
	fmt.Println(b)
	b %= 3
	fmt.Println(b)
}
