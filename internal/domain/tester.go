package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

// TestStatus represents a test status.
type TestStatus struct {
	Status string
	Reason string
}

// SynchronousTest represents a synchronous test.
type SynchronousTest struct {
	ID          uuid.UUID
	Name        string
	Command     string
	ScheduledAt time.Time
	LastRun     time.Time
	Status      TestStatus
}

// PeriodicTest represents a periodical test.
type PeriodicTest struct {
	ID       uuid.UUID
	Name     string
	Command  string
	Interval time.Duration
	LastRun  time.Time
	Status   TestStatus
}

// TestResult represents a test result.
type TestResult struct {
	ID        uuid.UUID
	Status    string
	Data      []byte
	Error     string
	StartTime time.Time
	EndTime   time.Time
}
