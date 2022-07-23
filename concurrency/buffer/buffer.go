package main

import "fmt"

func routine(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Done!")
}

func main() {
	ch := make(chan int, 3)

	go routine(ch)

	fmt.Println(<-ch)
}
