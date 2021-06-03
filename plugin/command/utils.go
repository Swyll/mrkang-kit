package command

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"

	"github.com/pkg/errors"
)

func ExecCommand(command, stdin string, stdout io.Writer, cmdch chan<- *exec.Cmd) error {
	if command == "" {
		return errors.New("Command is nil")
	}

	command = strings.Trim(command, " ")
	cms := strings.Split(command, " ")

	cmd := exec.Command(cms[0], cms[1:]...)
	if cmdch != nil {
		cmdch <- cmd
	}
	if stdout != nil {
		cmd.Stdout = stdout
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if stdin != "" {
		cmd.Stdin = bytes.NewReader([]byte(stdin))
	}

	err := cmd.Start()
	if err != nil {
		return errors.WithStack(err)
	}
	err = cmd.Wait()
	if err != nil {
		return errors.WithStack(err)
	}

	errStr := stderr.String()
	if errStr != "" {
		return errors.New(fmt.Sprintf("Command:%s error-->%s", command, errStr))
	}

	return nil
}

func kill(signal int) error {
	return exec.Command("kill", "-9", strconv.Itoa(signal)).Run()
}
