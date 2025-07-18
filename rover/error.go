package rover

import (
	"errors"
	"fmt"
)

type ErrInvalidCommand struct {
	str string
}

func (e ErrInvalidCommand) Error() string {
	return fmt.Sprintf("invalid command %q", e.str)
}

type ErrInvalidDirection struct {
	str string
}

func (e ErrInvalidDirection) Error() string {
	return fmt.Sprintf("invalid direction %q", e.str)
}

type ErrCommandNotConfigured struct {
	cmd Command
}

func (e ErrCommandNotConfigured) Error() string {
	return fmt.Sprintf("command not configured %q", e.cmd.String())
}

var ErrNoPlateuProvided = errors.New("no plateu provided")
