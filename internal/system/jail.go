package system

import "context"

type Jail struct {
	Name   string
	Status string
}

type JailSystem interface {
	List(ctx context.Context) ([]Jail, error)
}
