package main

import (
	"context"
	"flag"
	"log"
	"net"
	"path/filepath"

	"github.com/4rgetlahm/backports/backportRequest"
	"github.com/4rgetlahm/backports/tracker/launcher"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

var clientset *kubernetes.Clientset

type backportRequestServer struct {
	backportRequest.UnimplementedBackportRequestServiceServer
}

func (s *backportRequestServer) RunBackport(ctx context.Context, req *backportRequest.BackportRequest) (*backportRequest.BackportResponse, error) {
	if req.BaseBranch == "" || req.TargetBranch == "" {
		return &backportRequest.BackportResponse{}, status.Error(400, "Base branch and target branch are required")
	}
	if req.Commits == nil || len(req.Commits) == 0 {
		return &backportRequest.BackportResponse{}, status.Error(400, "Commits are required")
	}
	if req.Image == "" {
		return &backportRequest.BackportResponse{}, status.Error(400, "Image is required")
	}
	if req.Reference == "" {
		return &backportRequest.BackportResponse{}, status.Error(400, "Reference ObjectID is required")
	}

	jobName := launcher.LaunchBackportJob(clientset, req.Image, req.Reference, req.BaseBranch, req.TargetBranch, req.Commits)

	return &backportRequest.BackportResponse{
		JobName: jobName,
	}, nil
}

func initKubernetesClient() *kubernetes.Clientset {

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}

func main() {
	clientset = initKubernetesClient()
	lis, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	backportRequest.RegisterBackportRequestServiceServer(grpcServer, &backportRequestServer{})

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server %v", err)
	}
}
