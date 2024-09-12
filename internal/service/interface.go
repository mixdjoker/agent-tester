package service

import (
	"context"

	"github.com/gofrs/uuid"
)

// TesterService provides methods to interact with the tester service.
type TesterService interface {
	// Update updates the tester service and returns the number of records updated.
	Update(ctx context.Context) (int64, error)
	// UpdateByID updates a specific record in the tester service by its ID.
	UpdateByID(ctx context.Context, id uuid.UUID) error
	// GetResults retrieves the results from the tester service.
	GetResults(ctx context.Context) ([]byte, error)
}
