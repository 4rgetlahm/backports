package service

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/4rgetlahm/backports/api/database"
	"github.com/4rgetlahm/backports/api/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRepository(owner string, name string) (entity.Repository, error) {
	var repository entity.Repository

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := database.RepositoryCollection.FindOne(ctx, bson.M{"owner": owner, "name": name}).Decode(&repository)

	if err != nil {
		return entity.Repository{}, errors.New("error retrieving repository")
	}

	return repository, nil
}

func CreateRepository(clone_url string, image string) (entity.Repository, error) {

	regex := regexp.MustCompile(`(?P<Server>..+)\/(?P<Owner>..+)\/(?P<Name>..+).git`)
	match := regex.FindStringSubmatch(clone_url)

	if len(match) != 4 {
		return entity.Repository{}, errors.New("invalid clone URL")
	}

	server := match[1]
	owner := match[2]
	name := match[3]

	existingRepository, err := GetRepository(owner, name)

	if err == nil {
		return existingRepository, errors.New("repository already exists")
	}

	repository := entity.Repository{
		ID:       primitive.NewObjectID(),
		Server:   server,
		Owner:    owner,
		Name:     name,
		CloneURL: clone_url,
		Image:    image,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var result *mongo.InsertOneResult
	result, err = database.RepositoryCollection.InsertOne(ctx, repository)

	if err != nil {
		return entity.Repository{}, errors.New("error creating repository")
	}

	repository.ID = result.InsertedID.(primitive.ObjectID)

	return repository, nil
}

func GetRepositories() ([]entity.Repository, error) {
	var repositories []entity.Repository

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cursor, err := database.RepositoryCollection.Find(ctx, bson.M{})

	if err != nil {
		return []entity.Repository{}, errors.New("Error retrieving repositories")
	}

	if err = cursor.All(ctx, &repositories); err != nil {
		return []entity.Repository{}, errors.New("Error retrieving repositories")
	}

	return repositories, nil
}
