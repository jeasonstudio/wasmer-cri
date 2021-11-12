package wasmercri

import (
	"context"

	"github.com/jeasonstudio/wasmer-cri/os"
	log "github.com/sirupsen/logrus"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// RuntimeServer is used to implement images
type RuntimeServer struct {
	pb.UnimplementedRuntimeServiceServer
	sanboxStore *PodSandboxStore
	os          *os.RealOS
}

// NewRuntimeServer register image server
func NewRuntimeServer() (*RuntimeServer, error) {
	return &RuntimeServer{
		sanboxStore: NewPodSandboxStore(),
		os:          &os.RealOS{},
	}, nil
}

// StopPodSandbox pod
func (s *RuntimeServer) StopPodSandbox(ctx context.Context, in *pb.StopPodSandboxRequest) (*pb.StopPodSandboxResponse, error) {
	return nil, nil
}

// RemovePodSandbox pod
func (s *RuntimeServer) RemovePodSandbox(ctx context.Context, in *pb.RemovePodSandboxRequest) (*pb.RemovePodSandboxResponse, error) {
	return nil, nil
}

// PodSandboxStatus pod
func (s *RuntimeServer) PodSandboxStatus(ctx context.Context, in *pb.PodSandboxStatusRequest) (*pb.PodSandboxStatusResponse, error) {
	return nil, nil
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

// ReopenContainerLog container
func (s *RuntimeServer) ReopenContainerLog(ctx context.Context, req *pb.ReopenContainerLogRequest) (*pb.ReopenContainerLogResponse, error) {
	return nil, nil
}

// ExecSync container
func (s *RuntimeServer) ExecSync(ctx context.Context, req *pb.ExecSyncRequest) (*pb.ExecSyncResponse, error) {
	return nil, nil
}

// Exec container
func (s *RuntimeServer) Exec(ctx context.Context, req *pb.ExecRequest) (*pb.ExecResponse, error) {
	return nil, nil
}

// Attach container
func (s *RuntimeServer) Attach(ctx context.Context, req *pb.AttachRequest) (*pb.AttachResponse, error) {
	return nil, nil
}

// PortForward container
func (s *RuntimeServer) PortForward(ctx context.Context, req *pb.PortForwardRequest) (*pb.PortForwardResponse, error) {
	return nil, nil
}

// ContainerStats container
func (s *RuntimeServer) ContainerStats(ctx context.Context, req *pb.ContainerStatsRequest) (*pb.ContainerStatsResponse, error) {
	return nil, nil
}

// ListContainerStats container
func (s *RuntimeServer) ListContainerStats(ctx context.Context, req *pb.ListContainerStatsRequest) (*pb.ListContainerStatsResponse, error) {
	return nil, nil
}

// UpdateRuntimeConfig container
func (s *RuntimeServer) UpdateRuntimeConfig(ctx context.Context, req *pb.UpdateRuntimeConfigRequest) (*pb.UpdateRuntimeConfigResponse, error) {
	return nil, nil
}

// Status container
func (s *RuntimeServer) Status(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	return nil, nil
}
