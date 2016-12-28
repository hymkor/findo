package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/dustin/go-humanize"
)

var flagfileOnly = flag.Bool("f", false, "Select fileonly(Remove directories")
var nameOnly = flag.Bool("1", false, "Show nameonly(No Size,timestamp)")

func Main(args []string) error {
	var rx *regexp.Regexp
	if len(args) >= 1 {
		var err error
		rx, err = regexp.Compile("(?i)" + args[0])
		if err != nil {
			return err
		}
	}

	filepath.Walk(".", func(path_ string, info_ os.FileInfo, err_ error) error {
		name := filepath.Base(path_)
		if name == "." || name == ".." {
			return nil
		}
		if *flagfileOnly && info_.IsDir() {
			return nil
		}
		if rx == nil || rx.MatchString(name) {
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
