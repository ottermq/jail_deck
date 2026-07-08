package system

import (
	"context"
	"fmt"
	"strings"
)

type Command struct {
	Name string
	Args []string
}

type CommandResult struct {
	Stdout   string
	Stderr   string
	ExitCode int
}

type CommandRunner interface {
	Run(ctx context.Context, cmd Command) (CommandResult, error)
}

type CommandError struct {
	Command string
	Args    []string
	Result  CommandResult
	Err     error
}

func (e *CommandError) Error() string {
	return fmt.Sprintf("'%s %v' failed: %v", e.Command, strings.Join(e.Args, " "), e.Err)
}

func (e *CommandError) Unwrap() error {
	return e.Err
}
