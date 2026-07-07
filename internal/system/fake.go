package system

import (
	"context"
	"fmt"

	"github.com/ottermq/jaildeck/internal/domain"
)

type FakeJailSystem struct {
	jails map[string]domain.Jail
}

func NewFakeJailSystem() *FakeJailSystem {
	return &FakeJailSystem{
		jails: map[string]domain.Jail{
			"nginx":    {Name: "nginx", Status: domain.JailStatusRunning},
			"postgres": {Name: "postgres", Status: domain.JailStatusStopped},
			"redis":    {Name: "redis", Status: domain.JailStatusRunning},
		},
	}
}

func (s *FakeJailSystem) List(ctx context.Context) ([]domain.Jail, error) {
	return []domain.Jail{
		s.jails["nginx"],
		s.jails["postgres"],
		s.jails["redis"],
	}, nil
}

func (s *FakeJailSystem) Start(ctx context.Context, name string) (domain.Jail, error) {
	jail, ok := s.jails[name]
	if !ok {
		return domain.Jail{}, fmt.Errorf("jail %q not found", name)
	}

	jail.Status = domain.JailStatusRunning
	s.jails[name] = jail

	return jail, nil
}

func (s *FakeJailSystem) Stop(ctx context.Context, name string) (domain.Jail, error) {
	jail, ok := s.jails[name]
	if !ok {
		return domain.Jail{}, fmt.Errorf("jail %q not found", name)
	}

	jail.Status = domain.JailStatusStopped
	s.jails[name] = jail

	return jail, nil
}

func (s *FakeJailSystem) Restart(ctx context.Context, name string) (domain.Jail, error) {
	jail, ok := s.jails[name]
	if !ok {
		return domain.Jail{}, fmt.Errorf("jail %q not found", name)
	}

	jail.Status = domain.JailStatusRunning
	s.jails[name] = jail

	return jail, nil
}
