package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func title(urls ...string) <-chan string {
	c := make(chan string)

	for _, url := range urls {
		go func(url string) {
			res, _ := http.Get(url)

			html, _ := ioutil.ReadAll(res.Body)

			r, _ := regexp.Compile("<title>(.*?)<\\/title>")

			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

func forward(origin <-chan string, dest chan string) {
	for {
		dest <- <-origin
	}
}

func multiplexer(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go forward(input1, c)
	go forward(input2, c)

	return c
}

func main() {
	c := multiplexer(
		title("https://google.com", "https://github.com"),
		title("https://amazon.com", "https://spotify.com"),
	)

	fmt.Println(<-c, <-c)
}
