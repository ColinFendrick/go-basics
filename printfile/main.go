package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Please provide exactly one exiting file to open.")
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
