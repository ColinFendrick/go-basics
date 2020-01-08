package main

import (
	"fmt"
	"io"
	"os"
)

type textReader struct{}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide exactly one file to open.")
		os.Exit(1)
	}

	tr := textReader{}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(tr, file)
}

func (textReader) Write(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}
