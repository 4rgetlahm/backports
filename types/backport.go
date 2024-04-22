package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Backport struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	Author           string             `json:"author"`
	Commits          []string           `json:"commits"`
	Repository       Repository         `json:"repository"`
	TargetBranchName string             `json:"targetBranchName"`
	NewBranchName    string             `json:"newBranchName"`
	Events           []BackportEvent    `json:"events"`
	DateCreated      time.Time          `json:"dateCreated"`
	DateUpdated      time.Time          `json:"dateUpdated"`
}
