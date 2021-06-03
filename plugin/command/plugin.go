package command

import (
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Command struct {
	startCommand  string
	stopCommand   string
	cannelCommand string

	stdin string
	cmdch chan *exec.Cmd
}

type opt func(*Command)

func NewCommandPlu(startCommand string, opts ...opt) *Command {
	command := &Command{
		startCommand: startCommand,
		cmdch:        make(chan *exec.Cmd, 1),
	}

	for _, opt := range opts {
		opt(command)
	}

	return command
}

func WithStopCommand(stopCommand string) opt {
	return func(c *Command) {
		c.stopCommand = stopCommand
	}
}

func WithCannelCommand(cannelCommand string) opt {
	return func(c *Command) {
		c.cannelCommand = cannelCommand
	}
}

func WithStin(stdin string) opt {
	return func(c *Command) {
		c.stdin = stdin
	}
}

func (c *Command) Run(stdout io.Writer) error {
	err := ExecCommand(c.startCommand, c.stdin, stdout, c.cmdch)
	if err != nil && !strings.Contains(err.Error(), "killed") {
		return errors.Wrap(err, "startCommand err:")
	}

	return nil
}

func (c *Command) Cannel() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	go func() {
		err := ExecCommand(c.stopCommand, "", nil, nil)
		if err != nil {
			fmt.Println(err)
			cancel()
		}
	}()

	select {
	case <-ctx.Done():
		select {
		case cmd := <-c.cmdch:
			if cmd.Process != nil {
				return kill(cmd.Process.Pid)
			}
		default:
			return nil
		}
	}

	return nil
}

func (c *Command) Stop() error {
	err := ExecCommand(c.stopCommand, "", nil, nil)
	if err != nil {
		return errors.Wrap(err, "stopCommand err:")
	}

	return nil
}