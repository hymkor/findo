package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
)

var flagfileOnly = flag.Bool("f", false, "Select fileonly(Remove directories")
var nameOnly = flag.Bool("1", false, "Show nameonly(No Size,timestamp)")

func Main(args []string) error {
	var pattern string
	if len(args) >= 1 {
		pattern = strings.ToUpper(args[0])
	}

	filepath.Walk(".", func(path_ string, info_ os.FileInfo, err_ error) error {
		name := filepath.Base(path_)
		if name == "." || name == ".." {
			return nil
		}
		if *flagfileOnly && info_.IsDir() {
			return nil
		}
		var matched bool
		if pattern == "" {
			matched = true
		} else {
			var err error
			matched, err = filepath.Match(pattern, strings.ToUpper(name))
			if err != nil {
				matched = false
			}
		}
		if matched {
			fmt.Println(path_)
			if !*nameOnly {
				fmt.Printf("%12s %s\n", humanize.Comma(info_.Size()), info_.ModTime().String())
			}
		}
		return nil
	})
	return nil
}

func main() {
	flag.Parse()
	if err := Main(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
