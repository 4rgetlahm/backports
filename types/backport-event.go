package types

const (
	ActionVirtualMachinePreparing = "VirtualMachinePreparing"
	ActionVirtualMachineError     = "VirtualMachineError"
	ActionVirtualMachineCreated   = "VirtualMachineCreated"
	ActionVirtualMachineExited    = "VirtualMachineExited"

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
