package main

import "fmt"

func mean(numbers ...float64) float64 {
	total := 0.0

	for _, n := range numbers {
		total += n
	}

	return total / float64(len(numbers))
}

func main() {
	fmt.Printf("Mean: %.2f\n", mean(1, 2, 3, 4, 45, 5))
}
