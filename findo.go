package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mattn/go-isatty"
)

var flagfileOnly = flag.Bool("f", false, "Select fileonly not including directories")
var quotation = flag.Bool("q", false, "Enclose filename with double-quotations")
var nameOnly = flag.Bool("1", false, "Show nameonly without size and timestamp")
var flagList = flag.Bool("l", false, "Show size and timestamp")
var startDir = flag.String("d", ".", "Set start Directory")
var execCmd = flag.String("x", "", "Execute a command replacing {} to FILENAME")
var in = flag.Duration("in", 0, "Files modified in the duration such as 300ms, -1.5h or 2h45m")
var ignoreDots = flag.Bool("ignoredots", false, "Ignore files and directory starting with dot")

func main1(args []string) error {

	patterns := make([]string, len(args))
	for i := 0; i < len(args); i++ {
		patterns[i] = strings.ToUpper(args[i])
	}

	rich := isatty.IsTerminal(os.Stdout.Fd())
	if *flagList {
		rich = true
	}
	if *nameOnly {
		rich = false
	}

	filepath.Walk(*startDir, func(path_ string, info_ os.FileInfo, err_ error) error {
		name := filepath.Base(path_)
		if name == "." || name == ".." {
			return nil
		}
		if *ignoreDots {
			if name[0] == '.' || path_[0] == '.' || strings.Contains(path_, string(os.PathSeparator)+".") {
				return nil
			}
		}
		if *flagfileOnly && info_.IsDir() {
			return nil
		}
		if len(patterns) > 0 {
			matched := false
			for _, pattern := range patterns {
				m, err := filepath.Match(pattern, strings.ToUpper(name))
				if err == nil && m {
					matched = true
					break
				}
			}
			if !matched {
				return nil
			}
		}
		if *in != 0 && time.Now().Sub(info_.ModTime()) > *in {
			return nil
		}

		if *quotation {
			path_ = `"` + path_ + `"`
		}
		if *execCmd != "" {
			system(strings.Replace(*execCmd, "{}", path_, -1))
		} else {
			fmt.Println(path_)
			if rich {
				fmt.Printf("%12s %s\n", humanize.Comma(info_.Size()), info_.ModTime().String())
			}
		}
		return nil
	})
	return nil
}

func main() {
	flag.Parse()
	if err := main1(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
