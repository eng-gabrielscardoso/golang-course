package main

import "fmt"

func main() {
	// Como while
	var i int

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Como for
	for j := 0; j <= 10; j++ {
		fmt.Println(j)
	}
}
