package launcher

import "go.mongodb.org/mongo-driver/bson/primitive"

var GlobalLauncher Launcher

type Launcher interface {
	LaunchBackportJob(volume string, vcs string, reference primitive.ObjectID, newBranchName string, targetBranchName string, commits []string) error
	LaunchVolumeGenerationJob(volumeName string, vcs string, cloneUrl string, overwrite bool) (string, error)
}
