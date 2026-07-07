package system

import (
	"context"

	"github.com/ottermq/jaildeck/internal/domain"
)

type JailSystem interface {
	List(ctx context.Context) ([]domain.Jail, error)
	Start(ctx context.Context, name string) (domain.Jail, error)
	Stop(ctx context.Context, name string) (domain.Jail, error)
	Restart(ctx context.Context, name string) (domain.Jail, error)
}
