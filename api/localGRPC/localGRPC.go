package localGRPC

import (
	"github.com/4rgetlahm/backports/backportRequest"
	"github.com/4rgetlahm/backports/repositoryVolumeGenerator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var VolumeGenerationClient repositoryVolumeGenerator.RepositoryVolumeGenerationServiceClient
var BackportRequestClient backportRequest.BackportRequestServiceClient

func Init() *grpc.ClientConn {
	conn, err := grpc.NewClient("localhost:5001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	VolumeGenerationClient = repositoryVolumeGenerator.NewRepositoryVolumeGenerationServiceClient(conn)
	BackportRequestClient = backportRequest.NewBackportRequestServiceClient(conn)

	return conn
}
