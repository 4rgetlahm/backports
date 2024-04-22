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
	TargetBranchName string             `json:"target_branch_name"`
	NewBranchName    string             `json:"new_branch_name"`
	Events           []BackportEvent    `json:"events"`
	DateCreated      time.Time          `json:"date_created"`
}
