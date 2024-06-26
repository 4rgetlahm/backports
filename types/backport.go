package types

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Backport struct {
	ID               primitive.ObjectID `bson:"_id" json:"id"`
	Author           string             `json:"author" bson:"author"`
	Commits          []string           `json:"commits" bson:"commits"`
	Repository       Repository         `json:"repository" bson:"repository"`
	TargetBranchName string             `json:"target_branch_name" bson:"target_branch_name"`
	NewBranchName    string             `json:"new_branch_name" bson:"new_branch_name"`
	Events           []BackportEvent    `json:"events" bson:"events"`
	DateCreated      time.Time          `json:"date_created" bson:"date_created"`
}
