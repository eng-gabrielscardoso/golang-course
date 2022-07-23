package main

import (
	"fmt"
	"time"
)

func talk(person, text string, qtd int) {
	for i := 0; i < qtd; i++ {
		time.Sleep(time.Second)

		fmt.Printf("%s: %s (iteration %d)\n", person, text, i+1)
	}
}

func main() {
	// talk("Mary", "Why have you do not talk to me?", 3)
	// talk("John", "I just could talk after you", 1)

	go talk("Mary", "Hey", 10)
	talk("John", "Ho", 5)

	fmt.Println("End")
}
