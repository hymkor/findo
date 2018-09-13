package main

import (
	"os"
	"os/exec"
)

func system(cmdline string) error {
	shell := os.Getenv("SHELL")
	cmd1 := exec.Command(shell, "-c", cmdline)
	cmd1.Stdout = os.Stdout
	cmd1.Stderr = os.Stderr
	cmd1.Stdin = os.Stdin
	return cmd1.Run()
}
