package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	/*
		resp.Body is io.ReadCloser type, which is
		type ReadCloser interface {
			Reader
			Closer
		}

		ReadCloser interface uses two interfaces (Reader, Closer) inside of it
		The purpose of the Reader interface as a way to take different input sources
		and create a common point of contact to homogenize those inputs to a byte slice.
		The reason we pass a byte slice into the function, instead of being returned a byte slice,
		is Read() is actually modifiying through pointers the original byte slice we give it.
		The return values (int, err) represent the number of bytes modified (int) and an error object if
		this fails. So we are not generally interested in the return of the Read() function.
		Reader is
		type Reader interface {
				Read([]byte) (int, error)
		}

		Closer is
		type Closer interface {
				Close() error
		}
	*/
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	/* This is the naive way to read the function
	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
	*/

	/* We can also use
	io.Copy(os.Stdout, resp.Body)
	This os.Stdout implements the Writer interface (has Write([]byte) (int, error))
	*/

	lw := logWriter{}

	/* func Copy(Writer, Reader) (int64, error)
	type Writer interface {
		Write([]byte) (int, error)
	}
	type Reader interface {
		Read([]byte) (int, error)
	}

	resp.Body is a reader type (see above, it has Read() fn), so we can use it in the second arg.
	We create the logWriter to implement the Writer type (use the Write() fn with signature) */
	io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
