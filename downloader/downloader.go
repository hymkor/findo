package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const BYTES_PER_DOT = 1024 * 1024

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s URL FILENAME\n", os.Args[0])
		return
	}
	res, err := http.Get(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[1], err.Error())
		return
	}
	defer res.Body.Close()
	w, err2 := os.Create(os.Args[2])
	if err2 != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", os.Args[2], err2.Error())
		return
	}
	defer w.Close()
	for {
		_, err3 := io.CopyN(w, res.Body, BYTES_PER_DOT)
		if err3 != nil {
			fmt.Fprint(os.Stderr, "\n")
			if err3 != io.EOF {
				fmt.Fprintln(os.Stderr, err3.Error())
			}
			return
		}
		fmt.Fprint(os.Stderr, ".")
	}
}
