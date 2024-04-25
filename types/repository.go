package types

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Server   string             `json:"server" bson:"server"`
	Owner    string             `json:"owner" bson:"owner"`
	Name     string             `json:"name" bson:"name"`
	CloneURL string             `json:"clone_url" bson:"clone_url"`
	Volume   Volume             `json:"volume" bson:"volume"`
}
