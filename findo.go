package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func Main() error {
	var rx *regexp.Regexp
	if len(os.Args) >= 2 {
		var err error
		rx, err = regexp.Compile(os.Args[1])
		if err != nil {
			return err
		}
	}

	filepath.Walk(".", func(path_ string, info_ os.FileInfo, err_ error) error {
		name := filepath.Base(path_)
		if name == "." || name == ".." {
			return nil
		}
		name = strings.ToLower(name)
		if rx == nil || rx.MatchString(name) {
			fmt.Println(path_)
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
