package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/4rgetlahm/backports/api/database"
	"github.com/4rgetlahm/backports/api/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBackports(from int, to int) ([]entity.Backport, error) {
	if from < 0 || to < 0 || from > to {
		return []entity.Backport{}, errors.New("invalid from and to parameters")
	}

	var backports []entity.Backport

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cursor, err := database.BackportCollection.Find(ctx, bson.M{}, options.Find().
		SetSort(bson.D{{Key: "dateCreated", Value: -1}}).
		SetSkip(int64(from)).
		SetLimit(int64(to-from)))

	if err != nil {
		log.Println(err)
		return []entity.Backport{}, errors.New("error retrieving backports")
	}

	if err = cursor.All(ctx, &backports); err != nil {
		log.Println(err)
		return []entity.Backport{}, errors.New("error retrieving backports")
	}

	if len(backports) == 0 {
		return []entity.Backport{}, errors.New("no backports found")
	}

	return backports, nil
}

func GetBackport(id primitive.ObjectID) (entity.Backport, error) {
	var backport entity.Backport

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := database.BackportCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&backport)

	if err != nil {
		return entity.Backport{}, errors.New("error retrieving backport")
	}

	return backport, nil
}

func CreateBackport(author string, commits []string, repositoryOwner string, repositoryName, targetBranch string) (entity.Backport, error) {
	repository, err := GetRepository(repositoryOwner, repositoryName)
	if err != nil {
		return entity.Backport{}, errors.New("error retrieving repository")
	}

	backport := entity.Backport{
		ID:           primitive.NewObjectID(),
		Author:       author,
		Commits:      commits,
		Repository:   repository,
		TargetBranch: targetBranch,
		Events:       []entity.BackportEvent{},
		DateCreated:  time.Now(),
		DateUpdated:  time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var result *mongo.InsertOneResult
	result, err = database.BackportCollection.InsertOne(ctx, backport)

	if err != nil {
		return entity.Backport{}, errors.New("error creating backport")
	}

	backport.ID = result.InsertedID.(primitive.ObjectID)

	return backport, nil
}
