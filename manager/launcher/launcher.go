package launcher

import "go.mongodb.org/mongo-driver/bson/primitive"

var GlobalLauncher Launcher

type Launcher interface {
	LaunchBackportJob(volume string, reference primitive.ObjectID, newBranchName string, targetBranchName string, commits []string) error
	LaunchVolumeGenerationJob(volumeName string, cloneUrl string, credentials string, overwrite bool) (string, error)
}
