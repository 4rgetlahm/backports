package service

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/4rgetlahm/backports/api/localGRPC"
	"github.com/4rgetlahm/backports/backportRequest"
	"github.com/4rgetlahm/backports/database"
	"github.com/4rgetlahm/backports/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetBackports(from int, to int) ([]types.Backport, error) {
	if from < 0 || to < 0 || from > to {
		return []types.Backport{}, errors.New("invalid from and to parameters")
	}

	var backports []types.Backport

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cursor, err := database.BackportCollection.Find(ctx, bson.M{}, options.Find().
		SetSort(bson.D{{Key: "dateCreated", Value: -1}}).
		SetSkip(int64(from)).
		SetLimit(int64(to-from)))

	if err != nil {
		log.Println(err)
		return []types.Backport{}, errors.New("error retrieving backports")
	}

	if err = cursor.All(ctx, &backports); err != nil {
		log.Println(err)
		return []types.Backport{}, errors.New("error retrieving backports")
	}

	if len(backports) == 0 {
		return []types.Backport{}, errors.New("no backports found")
	}

	return backports, nil
}

func GetBackport(id primitive.ObjectID) (types.Backport, error) {
	var backport types.Backport

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := database.BackportCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&backport)

	if err != nil {
		return types.Backport{}, errors.New("error retrieving backport")
	}

	return backport, nil
}

func CreateBackport(author string, commits []string, repositoryName, targetBranchName string, newBranchName string) (types.Backport, error) {

	repository, err := GetRepository(repositoryName)

	if err != nil {
		return types.Backport{}, errors.New("error retrieving repository")
	}

	backport := types.Backport{
		ID:               primitive.NewObjectID(),
		Author:           author,
		Commits:          commits,
		Repository:       repository,
		TargetBranchName: targetBranchName,
		NewBranchName:    newBranchName,
		Events:           []types.BackportEvent{},
		DateCreated:      time.Now().UTC(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var result *mongo.InsertOneResult
	result, err = database.BackportCollection.InsertOne(ctx, backport)

	if err != nil {
		return types.Backport{}, errors.New("error creating backport")
	}

	backport.ID = result.InsertedID.(primitive.ObjectID)

	localGRPC.BackportRequestClient.RunBackport(context.Background(), &backportRequest.BackportRequest{
		Reference:        backport.ID.Hex(),
		Volume:           repository.Volume.Name,
		Vcs:              repository.VersionControlSystem,
		NewBranchName:    newBranchName,
		TargetBranchName: targetBranchName,
		Commits:          backport.Commits,
	})

	return backport, nil
}

func AddBackportEvent(backportID primitive.ObjectID, event *types.BackportEvent) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := database.BackportCollection.UpdateOne(ctx, bson.M{"_id": backportID}, bson.M{"$push": bson.M{"events": event}})

	if err != nil {
		return errors.New("error adding event")
	}

	return nil
}
