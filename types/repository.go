package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
	ID                   primitive.ObjectID `bson:"_id" json:"id"`
	Name                 string             `json:"name" bson:"name"`
	CloneURL             string             `json:"clone_url" bson:"clone_url"`
	Volume               Volume             `json:"volume" bson:"volume"`
	VersionControlSystem string             `json:"version_control_system" bson:"version_control_system"`
}

const (
	VersionControlSystemGit       = "git"
	VersionControlSystemMercurial = "mercurial"
)
