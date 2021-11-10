package wasmercri

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// RuntimeServer is used to implement images
type RuntimeServer struct {
	pb.UnimplementedRuntimeServiceServer
	sanboxStore *PodSandboxStore
}

// NewRuntimeServer register image server
func NewRuntimeServer() (*RuntimeServer, error) {
	return &RuntimeServer{
		sanboxStore: NewPodSandboxStore(),
	}, nil
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

// RunPodSandbox run pod sandbox
func (s *RuntimeServer) RunPodSandbox(ctx context.Context, in *pb.RunPodSandboxRequest) (*pb.RunPodSandboxResponse, error) {
	log.WithFields(log.Fields{
		"config":         in.Config,
		"runtimeHandler": in.RuntimeHandler,
	}).Debug("RunPodSandbox")

	id := "POD." + strconv.Itoa(int(time.Now().UnixNano()))
	sandbox := NewSandbox(id, &pb.PodSandboxMetadata{
		Name:      id,
		Uid:       id,
		Namespace: "default",
	}, pb.PodSandboxState_SANDBOX_NOTREADY)

	if err := s.sanboxStore.Add(*sandbox); err != nil {
		return nil, errors.Wrapf(err, "failed to add sandbox %+v into store", sandbox)
	}
	return &pb.RunPodSandboxResponse{PodSandboxId: id}, nil
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

// ListPodSandbox pod
func (s *RuntimeServer) ListPodSandbox(ctx context.Context, in *pb.ListPodSandboxRequest) (*pb.ListPodSandboxResponse, error) {
	log.WithFields(log.Fields{
		"podId":    in.Filter.Id,
		"podState": in.Filter.State,
		"podLabel": in.Filter.LabelSelector,
	}).Debug("ListPodSandbox")

	sandboxesInStore := s.sanboxStore.List()
	var sandboxes []*pb.PodSandbox

	for _, sandboxInStore := range sandboxesInStore {
		sandboxes = append(sandboxes, sandboxInStore)
	}
	// TODO: filter
	// sandboxes = c.filterCRISandboxes(sandboxes, r.GetFilter())
	return &pb.ListPodSandboxResponse{Items: sandboxes}, nil
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
