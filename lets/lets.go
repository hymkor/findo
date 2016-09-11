package main

import (
	"fmt"
	"io"
	"os"
)

func main_(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("Usage: %s SUBCOMMAND...")
	}
	switch args[1] {
	case "download":
		return download(args[1:])
	default:
		return fmt.Errorf("%s %s: no such subcommand")
	}
}
func main() {
	if err := main_(os.Args); err != nil && err != io.EOF {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	} else {
		os.Exit(0)
	}
}
