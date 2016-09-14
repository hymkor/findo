package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const BYTES_PER_DOT = 1024 * 1024

func download(args []string) error {
	if len(args) < 2 {
		return errors.New("Usage: download URL FILENAME")
	}
	res, err := http.Get(args[0])
	if err != nil {
		return err
	}
	defer res.Body.Close()
	w, err2 := os.Create(args[1])
	if err2 != nil {
		return err2
	}
	defer w.Close()
	for {
		_, err3 := io.CopyN(w, res.Body, BYTES_PER_DOT)
		if err3 != nil {
			fmt.Fprint(os.Stderr, "\n")
			if err3 != io.EOF {
				return err3
			}
			return nil
		}
		fmt.Fprint(os.Stderr, ".")
	}
}
