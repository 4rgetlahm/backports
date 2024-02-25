package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Backport struct {
	ID           primitive.ObjectID `bson:"_id" json:"id"`
	Author       string             `json:"author"`
	Commits      []string           `json:"commits"`
	Repository   Repository         `json:"repository"`
	TargetBranch string             `json:"targetBranch"`
	History      []HistoryEvent     `json:"history"`
	DateCreated  string             `json:"dateCreated"`
	DateUpdated  string             `json:"dateUpdated"`
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

type HistoryEvent struct {
	Action      string `json:"action"`
	Content     string `json:"content"`
	Status      string `json:"status"`
	DateCreated string `json:"dateCreated"`
}
