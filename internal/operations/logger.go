package operations

import (
	"context"
	"time"
)

type Entry struct {
	Timestamp time.Time `json:"timestamp"`
	Operation string    `json:"operation"`
	Target    string    `json:"target"`
	Command   string    `json:"command"`
	ExitCode  int       `json:"exit_code"`
	Success   bool      `json:"success"`
	Error     string    `json:"error,omitempty"`
}

type Logger interface {
	Log(ctx context.Context, entry Entry) error
}
