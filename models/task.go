package models

import "time"

type Task struct {
	TaskID              string    `json:"task_id"`
	Description         string    `json:"description"`
	Status              string    `json:"status"`
	Points              int       `json:"points"`
	StartTime           time.Time `json:"start_time"`
	EndTime             time.Time `json:"end_time"`
	Verifiers           []string  `json:"verifiers"` // List of user IDs who can verify
	VerificationEnabled bool      `json:"verification_enabled"`
}
