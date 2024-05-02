package types

import "time"

const (
	ActionVirtualMachinePreparing = "VirtualMachinePreparing"

	ActionVolumeCreateStart   = "VolumeCreateStart"
	ActionVolumeCreateSuccess = "VolumeCreated"
	ActionVolumeCreateFailure = "VolumeCreateFailure"

	ActionVirtualMachineError   = "VirtualMachineError"
	ActionVirtualMachineCreated = "VirtualMachineCreated"
	ActionVirtualMachineExited  = "VirtualMachineExited"

	ActionRunnerStarted = "RunnerStarted"

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

	ActionGitCherryPickStart   = "GitCherryPickStart"
	ActionGitCherryPickSuccess = "GitCherryPickSuccess"
	ActionGitCherryPickFailure = "GitCherryPickFailure"

	ActionGitPushStart   = "GitPushStart"
	ActionGitPushSuccess = "GitPushSuccess"
	ActionGitPushFailure = "GitPushFailure"

	ActionRunnerExited = "RunnerExited"
)

type BackportEvent struct {
	Action      string                 `json:"action" bson:"action"`
	Content     map[string]interface{} `json:"content" bson:"content"`
	DateCreated time.Time              `json:"date_created" bson:"date_created"`
}
