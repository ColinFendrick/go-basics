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
			as an independent concurrent thread of control, or (child) goroutine,
			within the same address space.
		*/
		go checkLink(link, c)
	}
	/*
		The main goroutine quits before the child routines finish -- the main routine spawns children
		but does not wait for the children to complete. THis is why we need channels.
		We have to send the child information back to the main routine.
	*/
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // Anonymous function call
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

/*
CONCURRENCY vs PARALLELISM

Concurrency - multiple thread executing code. If one thread blocks, another is picked up.
This is essentially one core running multiple goroutines.

Parallelism - requires multiple cores. Multiple goroutines running at the same time
on a separate core.
*/
