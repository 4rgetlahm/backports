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

	ActionGitFetchStart   = "GitFetchStart"
	ActionGitFetchSuccess = "GitFetchSuccess"
	ActionGitFetchFailure = "GitFetchFailure"

	ActionGitCheckoutStart   = "GitCheckoutSuccess"
	ActionGitCheckoutSuccess = "GitCheckoutSuccess"
	ActionGitCheckoutFailure = "GitCheckoutFailure"

	ActionGitCheckoutNewBranchStart   = "GitCheckoutNewBranchStart"
	ActionGitCheckoutNewBranchSuccess = "GitCheckoutNewBranchSuccess"
	ActionGitCheckoutNewBranchFailure = "GitCheckoutNewBranchFailure"

	ActionGitPullStart   = "GitPullStart"
	ActionGitPullSuccess = "GitPullSuccess"
	ActionGitPullFailure = "GitPullFailure"

	ActionGitPushStart   = "GitPushStart"
	ActionGitPushSuccess = "GitPushSuccess"
	ActionGitPushFailure = "GitPushFailure"

	ActionGitCherryPickStart   = "GitCherryPickStart"
	ActionGitCherryPickSuccess = "GitCherryPickSuccess"
	ActionGitCherryPickFailure = "GitCherryPickFailure"
)

type BackportEvent struct {
	Action      string `json:"action"`
	Content     string `json:"content"`
	DateCreated string `json:"dateCreated"`
}
