package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://github.com",
	}

	c := make(chan string)

	for _, link := range links {
		go visit(link, c)
	}
	// waiting for channel to have some data
	for l := range c {
		// fmt.Println(<-c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			visit(link, c)
		}(l)
	}
}

func visit(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("error ", err)
		c <- "down"
		return
	}

	fmt.Println(link, "is up and running")
	c <- link
}
