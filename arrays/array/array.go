package main

import "fmt"

func main() {
	// são homogêneas e estáticas
	var notes [3]float64

	fmt.Println(notes)

	notes[0], notes[1], notes[2] = 7.0, 5.6, 8.9

	fmt.Println(notes)

	var total float64 = 0.0

	for i := 0; i < len(notes); i++ {
		total += notes[i]
	}

	var mean float64 = total / float64(len(notes))

	fmt.Printf("Mean: %f\n", mean)
}
