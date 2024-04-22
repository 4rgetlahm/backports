package types

import "time"

type Volume struct {
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	LastUpdated time.Time `json:"last_updated"`
}

const (
	VolumeStatusNotInitialized = "not_initialized"
	VolumeStatusCreating       = "creating"
	VolumeStatusReady          = "ready"
)
