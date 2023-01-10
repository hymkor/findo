package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/mattn/go-isatty"
)

var (
	flagfileOnly   = flag.Bool("f", false, "Select fileonly not including directories")
	flagQuotation  = flag.Bool("q", false, "Enclose filename with double-quotations")
	flagNameOnly   = flag.Bool("1", false, "Show nameonly without size and timestamp")
	flagList       = flag.Bool("l", false, "Show size and timestamp")
	flagStartDir   = flag.String("d", ".", "Set start Directory")
	flagExecCmd    = flag.String("x", "", "Execute a command replacing {} to FILENAME")
	flagExecWithQ  = flag.String("X", "", `Execute a command replacing {} to "FILENAME" (same as -x,-v and -q)`)
	flagIn         = flag.Duration("in", 0, "Files modified in the duration such as 300ms, -1.5h or 2h45m")
	flagNotIn      = flag.Duration("notin", 0, "Files modified not in the duration such as 300ms, -1.5h or 2h45m")
	flagIgnoreDots = flag.Bool("ignoredots", false, "Ignore files and directory starting with dot")
	flagVerbose    = flag.Bool("v", false, "verbose (use with -x)")
)

func eachfile(dirname string, walk func(string, os.FileInfo) error) {
	children, err := ioutil.ReadDir(dirname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", dirname, err)
		return
	}
	for _, child := range children {
		childpath := filepath.Join(dirname, child.Name())
		if err := walk(childpath, child); err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", childpath, err)
		}
	}
}

func main1(args []string) error {

	patterns := make([]string, len(args))
	for i := 0; i < len(args); i++ {
		patterns[i] = strings.ToUpper(args[i])
	}

	rich := isatty.IsTerminal(os.Stdout.Fd())
	if *flagList {
		rich = true
	}
	if *flagNameOnly {
		rich = false
	}

	var walk func(string, os.FileInfo) error
	walk = func(path string, info os.FileInfo) error {
		name := filepath.Base(path)
		if name == "." || name == ".." {
			return nil
		}
		if *flagIgnoreDots && name[0] == '.' {
			return nil
		}
		if info.IsDir() {
			eachfile(path, walk)
		}
		if *flagfileOnly && info.IsDir() {
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
		if *flagIn != 0 && time.Now().Sub(info.ModTime()) > *flagIn {
			return nil
		}
		if *flagNotIn != 0 && time.Now().Sub(info.ModTime()) <= *flagNotIn {
			return nil
		}
		if *flagExecWithQ != "" {
			*flagExecCmd = *flagExecWithQ
			*flagQuotation = true
			*flagVerbose = true
		}
		if *flagQuotation {
			path = `"` + path + `"`
		}
		if *flagExecCmd != "" {
			cmdline := strings.Replace(*flagExecCmd, "{}", path, -1)
			if *flagVerbose {
				fmt.Fprintln(os.Stderr, cmdline)
			}
			system(cmdline)
		} else {
			fmt.Println(path)
			if rich {
				fmt.Printf("%12s %s\n", humanize.Comma(info.Size()), info.ModTime().String())
			}
		}
		return nil
	}
	eachfile(*flagStartDir, walk)
	return nil
}

func main() {
	flag.Parse()
	if err := main1(flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
