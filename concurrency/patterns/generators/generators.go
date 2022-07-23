package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

// Google I/0 2012 Go Concurrency Patterns

// <-chan - onlyread channel

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

func main() {
	t1 := title("https://google.com", "https://youtube.com", "https://drive.google.com")
	t2 := title("https://github.com", "https://linkedin.com", "https://gitlab.com")

	fmt.Println("First: ", <-t1, "|", <-t2)
	fmt.Println("Second: ", <-t1, "|", <-t2)
	fmt.Println("Third: ", <-t1, "|", <-t2)
}
