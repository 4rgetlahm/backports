package grpcService

import (
	"context"
	"log"

	"github.com/4rgetlahm/backports/backportRequest"
	"github.com/4rgetlahm/backports/manager/launcher"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type BackportRequestServer struct {
	backportRequest.UnimplementedBackportRequestServiceServer
}

func (s *BackportRequestServer) RunBackport(ctx context.Context, req *backportRequest.BackportRequest) (*emptypb.Empty, error) {
	if req.NewBranchName == "" || req.TargetBranchName == "" {
		return nil, status.Error(codes.InvalidArgument, "Base branch and target branch are required")
	}
	if req.Commits == nil || len(req.Commits) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Commits are required")
	}
	if req.Volume == "" {
		return nil, status.Error(codes.InvalidArgument, "Volume is required")
	}
	if req.Reference == "" {
		return nil, status.Error(codes.InvalidArgument, "Reference ObjectID is required")
	}

	objId, err := primitive.ObjectIDFromHex(req.Reference)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid ObjectID")
	}

	err = launcher.GlobalLauncher.LaunchBackportJob(req.Volume, objId, req.NewBranchName, req.TargetBranchName, req.Commits)

	if err != nil {
		log.Printf("Failed to launch backport job: %v", err)
		return nil, status.Error(codes.Internal, "Failed to launch backport job")
	}

	return nil, nil
}
