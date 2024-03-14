package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Backport struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Author       string             `json:"author"`
	Commits      []string           `json:"commits"`
	Repository   Repository         `json:"repository"`
	TargetBranch string             `json:"targetBranch"`
	Events       []BackportEvent    `json:"events"`
	DateCreated  time.Time          `json:"dateCreated"`
	DateUpdated  time.Time          `json:"dateUpdated"`
}

const (
	ActionVirtualMachineCreated = "VirtualMachineCreated"
	ActionVirtualMachineExited  = "VirtualMachineExited"
	ActionGitFetch              = "GitFetch"
	ActionGitCheckout           = "GitCheckout"
	ActionGitPull               = "GitPull"
	ActionGitPush               = "GitPush"
	ActionGitCherryPick         = "GitCherryPick"
	ActionGitCheckoutNewBranch  = "GitCheckoutNewBranch"
)

const (
	Success = "Success"
	Failure = "Failure"
)

type BackportEvent struct {
	Action      string `json:"action"`
	Content     string `json:"content"`
	Status      string `json:"status"`
	DateCreated string `json:"dateCreated"`
}
