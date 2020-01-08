package main

import (
	"fmt"
	"io"
	"os"
)

type textWriter struct{}

func main() {
	// You have to provide the filename as the second arg to the cli
	if len(os.Args) != 2 {
		fmt.Println("Please provide exactly one file to open.")
		os.Exit(1)
	}

	tw := textWriter{}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(tw, file)
}

// Create the Writer implementation and have textWriter receive it
func (textWriter) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}
