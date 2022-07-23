package main

import (
	"fmt"
	"time"
)

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func primes(n int, c chan int) {
	begin := 2

	for i := 0; i < n; i++ {
		for prime := begin; ; prime++ {
			if isPrime(prime) {
				c <- prime

				begin = prime + 1

				time.Sleep(time.Millisecond * 100)

				break
			}
		}
	}

	close(c)
}

func main() {
	c := make(chan int, 50)

	go primes(cap(c), c)

	for prime := range c {
		fmt.Printf("%d \n", prime)
	}

	fmt.Println("End")
}
