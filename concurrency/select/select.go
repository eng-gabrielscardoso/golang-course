package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
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

func faster(url1, url2, url3 string) string {
	c1 := title(url1)
	c2 := title(url2)
	c3 := title(url3)

	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "All lose"
		// default:
		// 	return "No response received"
	}
}

func main() {
	winner := faster(
		"https://google.com", "https://github.com", "https://amazon.com",
	)

	fmt.Println(winner)
}
