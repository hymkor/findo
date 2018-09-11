package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/mattn/go-isatty"
)

var flagfileOnly = flag.Bool("f", false, "Select fileonly not including directories")
var quotation = flag.Bool("q", false, "Enclose filename with double-quotations")
var nameOnly = flag.Bool("1", false, "Show nameonly without size and timestamp")
var flagList = flag.Bool("l", false, "Show size and timestamp")
var startDir = flag.String("d", ".", "Set start Directory")
var execCmd = flag.String("x", "", "Execute a command replacing {} to FILENAME")

func system(cmdline string) error {
	const CMDVAR = "CMDVAR"

	orgcmdarg := os.Getenv(CMDVAR)
	defer os.Setenv(CMDVAR, orgcmdarg)

	os.Setenv(CMDVAR, cmdline)

	cmd1 := exec.Command("cmd.exe", "/c", "%"+CMDVAR+"%")
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	cmd1.Stdin = os.Stdin
	return cmd1.Run()
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
	if *nameOnly {
		rich = false
	}

	filepath.Walk(*startDir, func(path_ string, info_ os.FileInfo, err_ error) error {
		name := filepath.Base(path_)
		if name == "." || name == ".." {
			return nil
		}
		if *flagfileOnly && info_.IsDir() {
			return nil
		}
		var matched bool
		if len(patterns) <= 0 {
			matched = true
		} else {
			matched = false
			for _, pattern := range patterns {
				m, err := filepath.Match(pattern, strings.ToUpper(name))
				if err == nil && m {
					matched = true
				}
			}
		}
		if matched {
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
