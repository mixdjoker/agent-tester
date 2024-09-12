package task

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/mixdjoker/agent-tester/internal/logger"
	"github.com/mixdjoker/agent-tester/internal/service"
)

type serv struct{}

// NewService ...
func NewService() service.TesterService {
	return &serv{}
}

// Update ...
func (s *serv) Update(ctx context.Context) (int64, error) {
	logger.Error(ctx).Msg("Update Unimplemented")
	return 0, errors.New("Unimplemented")
}

// UpdateByID ...
func (s *serv) UpdateByID(ctx context.Context, id uuid.UUID) error {
	logger.Error(ctx).Msgf("UpdateByID %s Unimplemented", id.String())
	return errors.New("Unimplemented")
}

// GetResults ...
func (s *serv) GetResults(ctx context.Context) ([]byte, error) {
	logger.Error(ctx).Msg("GetResults Unimplemented")
	return []byte{}, errors.New("Unimplemented")
}
