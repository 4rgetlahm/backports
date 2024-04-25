package main

import (
	"log"
	"net"
	"os"

	"github.com/4rgetlahm/backports/backportRequest"
	"github.com/4rgetlahm/backports/database"
	"github.com/4rgetlahm/backports/manager/grpcService"
	"github.com/4rgetlahm/backports/manager/launcher"
	"github.com/4rgetlahm/backports/manager/tracker"
	"github.com/4rgetlahm/backports/repositoryVolumeGenerator"
	"google.golang.org/grpc"
)

func main() {

	var pubsubProject = "backport-automation"
	var pubsubTopic = "backport.runner.updates"
	var pubsubCredentials = ""

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("Usage: %s launcher type", os.Args[0])
	}

	if args[0] == "docker" {
		log.Default().Println("Using Docker Launcher")
		dockerLauncher := launcher.DockerLauncher{}
		dockerLauncher.InitClient(pubsubProject, pubsubTopic, pubsubCredentials)
		launcher.GlobalLauncher = dockerLauncher

	} else {
		log.Fatalf("Invalid launcher type")
	}

	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	database.Init()

	go tracker.StartStatusTracker(pubsubProject, pubsubTopic, pubsubCredentials)

	grpcServer := grpc.NewServer()
	backportRequest.RegisterBackportRequestServiceServer(grpcServer, &grpcService.BackportRequestServer{})
	repositoryVolumeGenerator.RegisterRepositoryVolumeGenerationServiceServer(grpcServer, &grpcService.RepositoryVolumeGeneratorServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server %v", err)
	}
}
