package service

import (
	"context"
	"errors"
	"time"

	"github.com/4rgetlahm/backports/api/localGRPC"
	"github.com/4rgetlahm/backports/database"
	"github.com/4rgetlahm/backports/repositoryVolumeGenerator"
	"github.com/4rgetlahm/backports/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRepository(name string) (types.Repository, error) {
	var repo types.Repository

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := database.RepositoryCollection.FindOne(ctx, bson.M{"name": name}).Decode(&repo)

	if err != nil {
		return types.Repository{}, errors.New("error retrieving repository")
	}

	return repo, nil
}

func CreateRepository(vcs string, cloneURL string, name string) (types.Repository, error) {

	if vcs != types.VersionControlSystemGit && vcs != types.VersionControlSystemMercurial {
		return types.Repository{}, errors.New("invalid version control system")
	}

	existingRepository, err := GetRepository(name)

	if err == nil {
		return existingRepository, errors.New("repository already exists")
	}

	var volumeName string = name + ".repo"

	repo := types.Repository{
		ID:                   primitive.NewObjectID(),
		VersionControlSystem: vcs,
		Name:                 name,
		CloneURL:             cloneURL,
		Volume: types.Volume{
			Name:        volumeName,
			Status:      types.VolumeStatusNotInitialized,
			LastUpdated: time.Now().UTC(),
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var result *mongo.InsertOneResult
	result, err = database.RepositoryCollection.InsertOne(ctx, repo)

	if err != nil {
		return types.Repository{}, errors.New("error creating repository")
	}

	repo.ID = result.InsertedID.(primitive.ObjectID)

	var trueBool = true

	localGRPC.VolumeGenerationClient.Generate(context.Background(), &repositoryVolumeGenerator.GenerateRepositoryVolumeRequest{
		VolumeName: volumeName,
		Vcs:        vcs,
		CloneUrl:   repo.CloneURL,
		Overwrite:  &trueBool,
	})

	return repo, nil
}

func GetRepositories() ([]types.Repository, error) {
	var repositories []types.Repository

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cursor, err := database.RepositoryCollection.Find(ctx, bson.M{})

	if err != nil {
		return []types.Repository{}, errors.New("error retrieving repositories")
	}

	if err = cursor.All(ctx, &repositories); err != nil {
		return []types.Repository{}, errors.New("error retrieving repositories")
	}

	return repositories, nil
}
