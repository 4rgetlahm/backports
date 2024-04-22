package service

import (
	"context"
	"time"

	"github.com/4rgetlahm/backports/database"
	"github.com/4rgetlahm/backports/types"
	"go.mongodb.org/mongo-driver/bson"
)

func GetReadyVolumes() ([]types.Volume, error) {
	var volumes []types.Volume

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	cursor, err := database.RepositoryCollection.Find(ctx, bson.M{
		"volume.status": bson.M{
			"$eq": types.VolumeStatusReady,
		},
	})

	if err != nil {
		return []types.Volume{}, err
	}

	if err = cursor.All(ctx, &volumes); err != nil {
		return []types.Volume{}, err
	}

	if len(volumes) == 0 {
		return []types.Volume{}, nil
	}

	return volumes, nil
}

func UpdateVolumeStatus(volumeName string, newStatus string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := database.RepositoryCollection.UpdateMany(ctx, bson.M{
		"volume.name": volumeName,
	}, bson.M{
		"$set": bson.M{
			"volume.status":      newStatus,
			"volume.lastupdated": time.Now(),
		},
	})

	if err != nil {
		return err
	}

	return nil
}
