// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.26.0
// source: repositoryVolumeGenerator.proto

package repositoryVolumeGenerator

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RepositoryVolumeGenerationServiceClient is the client API for RepositoryVolumeGenerationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RepositoryVolumeGenerationServiceClient interface {
	Generate(ctx context.Context, in *GenerateRepositoryVolumeRequest, opts ...grpc.CallOption) (*GenerateVolumeResponse, error)
}

type repositoryVolumeGenerationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRepositoryVolumeGenerationServiceClient(cc grpc.ClientConnInterface) RepositoryVolumeGenerationServiceClient {
	return &repositoryVolumeGenerationServiceClient{cc}
}

func (c *repositoryVolumeGenerationServiceClient) Generate(ctx context.Context, in *GenerateRepositoryVolumeRequest, opts ...grpc.CallOption) (*GenerateVolumeResponse, error) {
	out := new(GenerateVolumeResponse)
	err := c.cc.Invoke(ctx, "/RepositoryVolumeGenerationService/Generate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RepositoryVolumeGenerationServiceServer is the server API for RepositoryVolumeGenerationService service.
// All implementations must embed UnimplementedRepositoryVolumeGenerationServiceServer
// for forward compatibility
type RepositoryVolumeGenerationServiceServer interface {
	Generate(context.Context, *GenerateRepositoryVolumeRequest) (*GenerateVolumeResponse, error)
	mustEmbedUnimplementedRepositoryVolumeGenerationServiceServer()
}

// UnimplementedRepositoryVolumeGenerationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRepositoryVolumeGenerationServiceServer struct {
}

func (UnimplementedRepositoryVolumeGenerationServiceServer) Generate(context.Context, *GenerateRepositoryVolumeRequest) (*GenerateVolumeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Generate not implemented")
}
func (UnimplementedRepositoryVolumeGenerationServiceServer) mustEmbedUnimplementedRepositoryVolumeGenerationServiceServer() {
}

// UnsafeRepositoryVolumeGenerationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RepositoryVolumeGenerationServiceServer will
// result in compilation errors.
type UnsafeRepositoryVolumeGenerationServiceServer interface {
	mustEmbedUnimplementedRepositoryVolumeGenerationServiceServer()
}

func RegisterRepositoryVolumeGenerationServiceServer(s grpc.ServiceRegistrar, srv RepositoryVolumeGenerationServiceServer) {
	s.RegisterService(&RepositoryVolumeGenerationService_ServiceDesc, srv)
}

func _RepositoryVolumeGenerationService_Generate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateRepositoryVolumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RepositoryVolumeGenerationServiceServer).Generate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RepositoryVolumeGenerationService/Generate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RepositoryVolumeGenerationServiceServer).Generate(ctx, req.(*GenerateRepositoryVolumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RepositoryVolumeGenerationService_ServiceDesc is the grpc.ServiceDesc for RepositoryVolumeGenerationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RepositoryVolumeGenerationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RepositoryVolumeGenerationService",
	HandlerType: (*RepositoryVolumeGenerationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Generate",
			Handler:    _RepositoryVolumeGenerationService_Generate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "repositoryVolumeGenerator.proto",
}