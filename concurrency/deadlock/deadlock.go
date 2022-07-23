package main

import (
	"fmt"
	"time"
)

func routine(c chan int) {
	time.Sleep(time.Second)

	c <- 1

	fmt.Println("After read")
}

func main() {
	ch := make(chan int) // channel without buffer

	go routine(ch)

	fmt.Println(<-ch) // block operation
	fmt.Println("Readed")
	fmt.Println(<-ch)
}
