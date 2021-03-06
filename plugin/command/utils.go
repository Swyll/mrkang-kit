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

func ExecCommand(command, stdin string, stdout io.Writer, cmdch chan<- *exec.Cmd, envs []string) error {
	if command == "" {
		return errors.New("Command is nil")
	}

	command = strings.Trim(command, " ")

	cmd := exec.Command("bash", "-c", command)
	if cmdch != nil {
		cmdch <- cmd
	}
	if stdout != nil {
		cmd.Stdout = stdout
	}
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	if envs != nil {
		cmd.Env = envs
	}

	if stdin != "" {
		cmd.Stdin = bytes.NewReader([]byte(stdin))
	}

	err := cmd.Start()
	if err != nil {
		return errors.WithStack(err)
	}
	cmd.Wait()

	errStr := stderr.String()
	if errStr != "" {
		return errors.New(fmt.Sprintf("Command:%s error-->%s", command, errStr))
	}

	return nil
}

func kill(signal int) error {
	return exec.Command("kill", "-9", strconv.Itoa(signal)).Run()
}
