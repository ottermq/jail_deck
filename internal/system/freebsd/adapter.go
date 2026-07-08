package freebsd

import "github.com/ottermq/jaildeck/internal/system"

type Adapter struct {
	runner system.CommandRunner
}

func NewAdapter(runner system.CommandRunner) *Adapter {
	return &Adapter{
		runner: runner,
	}
}
