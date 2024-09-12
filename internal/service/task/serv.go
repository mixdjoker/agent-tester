package task

import (
	"context"
	"errors"

	"github.com/gofrs/uuid"
	"github.com/mixdjoker/agent-tester/internal/logger"
	"github.com/mixdjoker/agent-tester/internal/service"
)

type serv struct{}

// NewService initializes and returns a new instance of TesterService.
func NewService() service.TesterService {
	return &serv{}
}

// Update performs the update of all tests located in the 'tests' directory.
// It processes each test, applying necessary updates, and returns the total count of successfully updated tests.
// If any errors occur during the update process, it returns the error along with the count of updated tests up to that point.
func (s *serv) Update(ctx context.Context) (int64, error) {
	logger.Error(ctx).Msg("Update Unimplemented")
	return 0, errors.New("Unimplemented")
}

// UpdateByID performs the update of a specific test identified by the provided UUID.
// It locates the test in the 'tests' directory, applies the necessary updates, and returns an error if any issues occur during the process.
// If the update is successful, the method returns nil.
func (s *serv) UpdateByID(ctx context.Context, id uuid.UUID) error {
	logger.Error(ctx).Msgf("UpdateByID %s Unimplemented", id.String())
	return errors.New("Unimplemented")
}

// GetResults retrieves the results of the tests and returns them as a JSON-encoded byte slice.
// If an error occurs during the retrieval or encoding of the results, it returns the error.
// On success, it provides the JSON data representing the test results.
func (s *serv) GetResults(ctx context.Context) ([]byte, error) {
	logger.Error(ctx).Msg("GetResults Unimplemented")
	return []byte{}, errors.New("Unimplemented")
}
