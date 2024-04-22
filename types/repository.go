package types

import "go.mongodb.org/mongo-driver/bson/primitive"

type Repository struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Server   string             `json:"server"`
	Owner    string             `json:"owner"`
	Name     string             `json:"name"`
	CloneURL string             `json:"clone_url"`
}
