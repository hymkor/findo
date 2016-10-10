package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

type History struct {
	XMLName xml.Name `xml:"history"`
	Line    []string `xml:"line"`
}

func main() {
	history := History{Line: []string{"ahaha", "ihihi", "ufufu"}}

	data, err := xml.MarshalIndent(&history, "", " ")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
	fmt.Print(xml.Header)
	os.Stdout.Write(data)
}
