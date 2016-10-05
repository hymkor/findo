package main

import "syscall"

var kernel32 = syscall.NewLazyDLL("Kernel32")

var localAlloc = kernel32.NewProc("LocalAlloc")

func LocalAlloc(uflag uintptr, bytes uintptr) (uintptr, error) {
	rc, _, err := localAlloc.Call(uflag, bytes)
	if rc == 0 {
		return 0, err
	} else {
		return rc, nil
	}
}

var localFree = kernel32.NewProc("LocalFree")

func LocalFree(hMem uintptr) error {
	rc, _, err := localFree.Call(hMem)
	if rc == 0 {
		return nil
	} else {
		return err
	}
}
