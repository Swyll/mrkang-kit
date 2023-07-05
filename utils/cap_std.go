package utils

import (
	"os"
	"syscall"
)

var golab *os.File

func CapPanic(file string) error {
	golab, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}

	return syscall.Dup2(int(golab.Fd()), int(os.Stderr.Fd()))
}
