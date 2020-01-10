package main

import (
	"fmt"
	"net/http"
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

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- "Might be down I think"
		return
	}

	fmt.Println(link, "is up")
	c <- "Yep its up"
}

/*
Why doesn't this work?
The main routine spawns all the goroutines based off go checkLink().
After all those are spawned, it receives a blocking call with fmt.Println(<-c)
and goes to sleep while it awaits some information to be put into channel c.
Once some data is put into c, it executes that line of code then quits.
So the only thing that gets put into the channel is the first call to resolve.
*/
