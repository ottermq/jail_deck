package system

import "context"

type FakeJailSystem struct{}

func NewFakeJailSystem() *FakeJailSystem {
	return &FakeJailSystem{}
}

func (s *FakeJailSystem) List(ctx context.Context) ([]Jail, error) {
	return []Jail{
		{Name: "nginx", Status: "running"},
		{Name: "postgres", Status: "stopped"},
		{Name: "redis", Status: "running"},
	}, nil
}
