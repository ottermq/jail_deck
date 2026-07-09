package services

import (
	"context"

	"github.com/ottermq/jaildeck/internal/operations"
)

type OperationService struct {
	reader operations.Reader
}

func NewOperationService(reader operations.Reader) *OperationService {
	return &OperationService{reader: reader}
}

func (s *OperationService) Recent(ctx context.Context, limit int, filters map[string]any) ([]operations.Entry, error) {
	return s.reader.Recent(ctx, limit, filters)
}
