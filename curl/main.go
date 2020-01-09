package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func init() {
	if len(os.Args) != 2 && len(os.Args) != 3 {
		fmt.Println("Usage: go run . <url> <?filename>")
		os.Exit(1)
	}
}

func main() {
	url := os.Args[1]
	name := "something.html"

	if len(os.Args) == 3 {
		name = os.Args[2]
	}

	resp, err := http.Get(url)
	check(err)
	defer resp.Body.Close()

	f, err := os.Create(name)
	check(err)

	io.Copy(f, resp.Body)
	fmt.Printf("Successfully wrote the contents of %v to %v.", url, name)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
