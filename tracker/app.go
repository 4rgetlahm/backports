package main

import (
	"context"
	"log"
	"net"

	"github.com/4rgetlahm/backports/backportRequest"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

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
	return &backportRequest.BackportResponse{
		Pod: "backport-1",
	}, nil
}

func initKubernetesClient() {
	rules := clientcmd.NewDefaultClientConfigLoadingRules()
	kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
	config, err := kubeconfig.ClientConfig()
	if err != nil {
		log.Fatalf("Failed to load kubeconfig: %v", err)
	}

	clientset := kubernetes.NewForConfigOrDie(config)

	nodeList, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("Failed to list nodes: %v", err)
	}

	for _, node := range nodeList.Items {
		log.Printf("Node: %s\n", node.Name)
	}
}

func main() {
	initKubernetesClient()
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
