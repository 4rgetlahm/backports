package types

import "time"

type Volume struct {
	Name        string    `json:"name" bson:"name"`
	Status      string    `json:"status" bson:"status"`
	LastUpdated time.Time `json:"last_updated" bson:"last_updated"`
}

const (
	VolumeStatusNotInitialized = "not_initialized"
	VolumeStatusCreating       = "creating"
	VolumeStatusReady          = "ready"
)
