package main

import (
	"os"
	"os/exec"
)

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
