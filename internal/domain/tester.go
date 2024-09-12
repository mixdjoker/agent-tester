package domain

import (
	"time"

	"github.com/gofrs/uuid"
)

type TestStatus struct {
	Status string
	Reason string
}

type SynchronoisTest struct {
	ID          uuid.UUID
	Name        string
	Command     string
	ScheduledAt time.Time
	LastRun     time.Time
	Status      TestStatus
}

type PeriodicTest struct {
	ID       uuid.UUID
	Name     string
	Command  string
	Interval time.Duration
	LastRun  time.Time
	Status   TestStatus
}

type TestResult struct {
	ID        uuid.UUID
	Status    string
	Data      []byte
	Error     string
	StartTime time.Time
	EndTime   time.Time
}
