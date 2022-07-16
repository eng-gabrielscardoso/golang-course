package main

import "fmt"

func factor(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("Error: %d", n)

	case n == 0:
		return 1, nil

	default:
		previousFactor, _ := factor(n - 1)
		return n * previousFactor, nil
	}
}

func main() {
	result, _ := factor(5)

	fmt.Println(result)

	_, error := factor(-4)

	fmt.Println(error)
}
