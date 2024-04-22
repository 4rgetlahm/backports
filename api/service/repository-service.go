package service

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/4rgetlahm/backports/api/localGRPC"
	"github.com/4rgetlahm/backports/database"
	"github.com/4rgetlahm/backports/repositoryVolumeGenerator"
	"github.com/4rgetlahm/backports/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRepository(owner string, name string) (types.Repository, error) {
	var repo types.Repository

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := database.RepositoryCollection.FindOne(ctx, bson.M{"owner": owner, "name": name}).Decode(&repo)

	if err != nil {
		return types.Repository{}, errors.New("error retrieving repository")
	}

	return repo, nil
}

func CreateRepository(clone_url string) (types.Repository, error) {

	regex := regexp.MustCompile(`(?P<Server>..+)\/(?P<Owner>..+)\/(?P<Name>..+).git`)
	match := regex.FindStringSubmatch(clone_url)

	if len(match) != 4 {
		return types.Repository{}, errors.New("invalid clone URL")
	}

	server := match[1]
	owner := match[2]
	name := match[3]

	existingRepository, err := GetRepository(owner, name)

	if err == nil {
		return existingRepository, errors.New("repository already exists")
	}

	var volumeName string = owner + "." + name

	repo := types.Repository{
		ID:       primitive.NewObjectID(),
		Server:   server,
		Owner:    owner,
		Name:     name,
		CloneURL: clone_url,
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
		VolumeName:  volumeName,
		CloneUrl:    repo.CloneURL,
		Credentials: "",
		Overwrite:   &trueBool,
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
