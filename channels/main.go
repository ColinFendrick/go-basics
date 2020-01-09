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
		"http://amazon.com",
		"asdflkasfdlksajfds.asldfjsldajflskdafjsdfalkj",
	}
	/*
		Create a new channel with make(chan val-type).
		Channels are typed by the values they convey.
	*/
	c := make(chan string)

	for _, link := range links {
		/*
			// A "go" statement starts the execution of a function call
			as an independent concurrent thread of control, or goroutine,
			within the same address space.
		*/
		go checkLink(link, c)
	}

	for l := range c {
		go func(link string) { // Anonymous function call
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
