package operations

import (
	"context"
	"time"
)

type Entry struct {
	Timesamp  time.Time `json:"timestamp"`
	Operation string    `json:"operation"`
	Target    string    `json:"target"`
	Commnad   string    `json:"command"`
	ExitCode  int       `json:"exit_code"`
	Success   bool      `json:"success"`
	Error     string    `json:"error,omitempty"`
}

type Logger interface {
	Log(ctx context.Context, entry Entry) error
}
