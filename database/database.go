package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var BackportCollection *mongo.Collection
var RepositoryCollection *mongo.Collection

func Init() {
	DB = GetDatabase()
	BackportCollection = DB.Collection("backports")
	RepositoryCollection = DB.Collection("repositories")
}

func GetDatabase() *mongo.Database {
	if DB == nil {
		err := Connect()
		if err != nil {
			panic(err)
		}
	}
	return DB
}

func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017/"))
	if err != nil {
		fmt.Println("Error connecting to MongoDB")
	}
	DB = client.Database("backport-automation")
	return err
}
