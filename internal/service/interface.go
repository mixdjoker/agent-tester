package service

import (
	"context"

	"github.com/gofrs/uuid"
)

type TesterService interface {
	Update(ctx context.Context) (int64, error)
	UpdateByID(ctx context.Context, id uuid.UUID) error
	GetResults(ctx context.Context) ([]byte, error)
}
