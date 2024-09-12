package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

// TestStatus ...
type TestStatus struct {
	Status string
	Reason string
}

// SynchronoisTest ...
type SynchronoisTest struct {
	ID          uuid.UUID
	Name        string
	Command     string
	ScheduledAt time.Time
	LastRun     time.Time
	Status      TestStatus
}

// PeriodicTest ...
type PeriodicTest struct {
	ID       uuid.UUID
	Name     string
	Command  string
	Interval time.Duration
	LastRun  time.Time
	Status   TestStatus
}

// TestResult ...
type TestResult struct {
	ID        uuid.UUID
	Status    string
	Data      []byte
	Error     string
	StartTime time.Time
	EndTime   time.Time
}
