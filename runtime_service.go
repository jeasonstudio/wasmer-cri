package wasmercri

import (
	"context"
	"log"

	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// RuntimeServer is used to implement images
type RuntimeServer struct {
	pb.UnimplementedRuntimeServiceServer
}

// NewRuntimeServer register image server
func NewRuntimeServer() (*RuntimeServer, error) {
	return &RuntimeServer{}, nil
}

// Version runtime version
func (s *RuntimeServer) Version(ctx context.Context, in *pb.VersionRequest) (*pb.VersionResponse, error) {
	log.Printf("Version Received: %v", in)
	return &pb.VersionResponse{
		Version:           "1.0.0",
		RuntimeName:       "wasmer",
		RuntimeVersion:    "1.0.0",
		RuntimeApiVersion: "1.0.0",
	}, nil
}

// CreateContainer create container
func (s *RuntimeServer) CreateContainer(ctx context.Context, in *pb.CreateContainerRequest) (*pb.CreateContainerResponse, error) {
	log.Printf("CreateContainer Received: %v", in)
	return &pb.CreateContainerResponse{}, nil
}

// StartContainer start container
func (s *RuntimeServer) StartContainer(ctx context.Context, in *pb.StartContainerRequest) (*pb.StartContainerResponse, error) {
	log.Printf("StartContainer Received: %v", in)
	return &pb.StartContainerResponse{}, nil
}

// StopContainer stop container
func (s *RuntimeServer) StopContainer(ctx context.Context, in *pb.StopContainerRequest) (*pb.StopContainerResponse, error) {
	log.Printf("StopContainer Received: %v", in)
	return &pb.StopContainerResponse{}, nil
}

// RemoveContainer remove container
func (s *RuntimeServer) RemoveContainer(ctx context.Context, in *pb.RemoveContainerRequest) (*pb.RemoveContainerResponse, error) {
	log.Printf("RemoveContainer Received: %v", in)
	return &pb.RemoveContainerResponse{}, nil
}

// ListContainers list containers
func (s *RuntimeServer) ListContainers(ctx context.Context, in *pb.ListContainersRequest) (*pb.ListContainersResponse, error) {
	log.Printf("ListContainers Received: %v", in)
	return &pb.ListContainersResponse{}, nil
}

// ContainerStatus show container status
func (s *RuntimeServer) ContainerStatus(ctx context.Context, in *pb.ContainerStatusRequest) (*pb.ContainerStatusResponse, error) {
	log.Printf("ContainerStatus Received: %v", in)
	return &pb.ContainerStatusResponse{}, nil
}

// UpdateContainerResources update container resources
func (s *RuntimeServer) UpdateContainerResources(ctx context.Context, in *pb.UpdateContainerResourcesRequest) (*pb.UpdateContainerResourcesResponse, error) {
	log.Printf("UpdateContainerResources Received: %v", in)
	return &pb.UpdateContainerResourcesResponse{}, nil
}
