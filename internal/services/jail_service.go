package services

import (
	"context"

	"github.com/otterlabs/jaildeck/internal/system"
)

type JailService struct {
	system system.JailSystem
}

func NewJailService(system system.JailSystem) *JailService {
	return &JailService{
		system: system,
	}
}

func (s *JailService) List(ctx context.Context) ([]system.Jail, error) {
	return s.system.List(ctx)
}
