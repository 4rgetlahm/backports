package grpcService

import (
	"context"
	"log"

	"github.com/4rgetlahm/backports/manager/launcher"
	"github.com/4rgetlahm/backports/repositoryVolumeGenerator"
	"google.golang.org/grpc/status"
)

type RepositoryVolumeGeneratorServer struct {
	repositoryVolumeGenerator.UnimplementedRepositoryVolumeGenerationServiceServer
}

func (s *RepositoryVolumeGeneratorServer) Generate(ctx context.Context, req *repositoryVolumeGenerator.GenerateRepositoryVolumeRequest) (*repositoryVolumeGenerator.GenerateVolumeResponse, error) {
	if req.CloneUrl == "" {
		return &repositoryVolumeGenerator.GenerateVolumeResponse{}, status.Error(400, "Clone URL is required")
	}

	if req.VolumeName == "" {
		return &repositoryVolumeGenerator.GenerateVolumeResponse{}, status.Error(400, "Volume name is required")
	}

	var overwrite bool
	if req.Overwrite == nil {
		overwrite = false
	} else {
		overwrite = *req.Overwrite
	}

	id, err := launcher.GlobalLauncher.LaunchVolumeGenerationJob(req.VolumeName, req.Vcs, req.CloneUrl, overwrite)

	if err != nil {
		log.Printf("Failed to launch volume generation job: %v", err)
		return &repositoryVolumeGenerator.GenerateVolumeResponse{}, status.Error(500, "Failed to launch volume generation job")
	}

	return &repositoryVolumeGenerator.GenerateVolumeResponse{
		VolumeName:  req.VolumeName,
		ContainerId: id,
	}, nil
}
