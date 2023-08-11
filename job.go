package jobber

import (
	"time"
)

type Job interface {
	// Get JobId
	Id() string

	// Get Job Metadata
	Metadata() map[string]string

	// Get Job Creation Time
	CreatedTime() time.Time

	// Get Job last Update Time
	UpdatedTime() time.Time

	// Get Current Retry attempt count. Attempt starts from 1
	CurrentAttemptCount() uint32

	// Get Max Retry Attempts count
	MaxAttemptCount() uint32

	// Update job metadata
	UpdateMetadata() error

	// Update LastActive Timestamp. Signals Worker is still active on job.
	// Otherwise job might get assigned to another worker after KeepAlive Expiry
	// This should be handled by Job Manager unless clients disable auto-keep-alive and run keep alive explicitly
	UpdateLastActiveTimestamp() error

	// Update Job status to COMPLETED
	SetCompleted() error

	// Update Job status to FAILED
	SetFailed() error

	// Update Job status to RETRY
	SetRetry() error
}
