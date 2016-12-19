package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/dustin/go-humanize"
)

func Main() error {
	var rx *regexp.Regexp
	if len(os.Args) >= 2 {
		var err error
		rx, err = regexp.Compile("(?i)" + os.Args[1])
		if err != nil {
			return err
		}
	}

	filepath.Walk(".", func(path_ string, info_ os.FileInfo, err_ error) error {
		name := filepath.Base(path_)
		if name == "." || name == ".." {
			return nil
		}
		if rx == nil || rx.MatchString(name) {
			fmt.Println(path_)
			fmt.Printf("%12s %s\n",humanize.Comma(info_.Size()),info_.ModTime().String())
		}
		return nil
	})
	return nil
}

func main() {
	if err := Main(); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
