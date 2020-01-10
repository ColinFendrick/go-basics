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
			A "go" statement starts the execution of a function call
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
	for l := range c { // Wait for channel to return a value, then run body. An infinite loop that runs every time channel emits a value.
		go func(link string) { // We don't want to pause the main routine, so we create a separate routine to house the throttling and the routines.
			time.Sleep(5 * time.Second)
			checkLink(link, c) // We have to call the goroutine fn with l, rather than referencing l directly, so we can reference different memory locations.
		}(l) // Function literal (anonymous function/lambda)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // Send in the link to the channel
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
