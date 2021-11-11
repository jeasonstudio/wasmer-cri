package wasmercri

import (
	"context"
	"path/filepath"
	"time"

	"github.com/jeasonstudio/wasmer-cri/os"
	"github.com/jeasonstudio/wasmer-cri/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

const (
	// sandboxesDir contains all sandbox root. A sandbox root is the running
	// directory of the sandbox, all files created for the sandbox will be
	// placed under this directory.
	sandboxesDir = "sandboxes"
	rootDir      = ".wasmerd"
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
func (s *RuntimeServer) RunPodSandbox(ctx context.Context, in *pb.RunPodSandboxRequest) (_ *pb.RunPodSandboxResponse, retErr error) {
	logger := log.WithContext(ctx)

	logger.WithFields(log.Fields{
		"config":         in.Config,
		"runtimeHandler": in.RuntimeHandler,
	}).Debug("RunPodSandbox")

	// Generate unique id and name for the sandbox and reserve the name.
	id := utils.GenerateID()
	logger.WithField("podsandboxid", id).Debug("generated id for pod sandbox")

	metadata := in.Config.GetMetadata()
	if metadata == nil {
		return nil, errors.New("Sandbox config must include metadata")
	}
	name := utils.MakeSandboxName(metadata)
	logger.WithField("podsanboxname", name).Debug("generated name for pod sandbox")

	timestamp := time.Now().UnixNano()

	sandbox := &pb.PodSandbox{
		Id:             id,
		Metadata:       metadata,
		State:          pb.PodSandboxState_SANDBOX_NOTREADY,
		CreatedAt:      timestamp,
		Labels:         map[string]string{},
		Annotations:    map[string]string{},
		RuntimeHandler: "wasmer",
	}
	logger.WithField("podsandbox", sandbox).Debug("create sandbox")

	var image interface{}
	logger.WithField("image", image).Debug("prepare image(ensure sandbox container image snapshot)")
	// TODO: Ensure sandbox container image snapshot.

	var runtime interface{}
	logger.WithField("runtime", runtime).Debug("prepare webassembly runtime(wasmer)")
	// TODO: init wasmer runtime

	// TODO: pod network (cni)

	// Create sandbox container root directories.
	sandboxRootDir := s.getSandboxRootDir(id)
	if err := s.os.MkdirAll(sandboxRootDir, 0755); err != nil {
		return nil, errors.Wrapf(err, "failed to create sandbox root directory %q",
			sandboxRootDir)
	}
	defer func() {
		if retErr != nil {
			// Cleanup the sandbox root directory.
			if err := s.os.RemoveAll(sandboxRootDir); err != nil {
				logger.WithError(err).Errorf("Failed to remove sandbox root directory %q",
					sandboxRootDir)
			}
		}
	}()
	logger.WithField("sandboxRootDir", sandboxRootDir).Debug("generate sandbox root dir")

	// TODO: setup sandbox files

	// TODO: update sandbox status such as pid
	sandbox.State = pb.PodSandboxState_SANDBOX_READY

	if err := s.sanboxStore.Add(sandbox); err != nil {
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
func (s *RuntimeServer) ListPodSandbox(ctx context.Context, in *pb.ListPodSandboxRequest) (_ *pb.ListPodSandboxResponse, _ error) {
	logger := log.WithContext((ctx))

	logger.WithFields(log.Fields{
		"podId":    in.Filter.Id,
		"podState": in.Filter.State,
		"podLabel": in.Filter.LabelSelector,
	}).Debug("ListPodSandbox")

	sandboxesInStore := s.sanboxStore.List()
	var sandboxes []*pb.PodSandbox

	for _, sb := range sandboxesInStore {
		// TODO: needs to filter sandboxes
		sandboxes = append(sandboxes, &pb.PodSandbox{
			Id:             sb.GetId(),
			Metadata:       sb.GetMetadata(),
			State:          sb.GetState(),
			CreatedAt:      sb.GetCreatedAt(),
			Labels:         sb.GetLabels(),
			Annotations:    sb.GetAnnotations(),
			RuntimeHandler: sb.GetRuntimeHandler(),
		})
	}

	return &pb.ListPodSandboxResponse{Items: sandboxes}, nil
}

// CreateContainer create container
func (s *RuntimeServer) CreateContainer(ctx context.Context, in *pb.CreateContainerRequest) (*pb.CreateContainerResponse, error) {
	logger := log.WithContext(ctx)
	logger.WithFields(log.Fields{
		"PodSandboxId":     in.PodSandboxId,
		"Image":            in.Config.Image.Image,
		"Command":          in.Config.Command,
		"SandboxHost":      in.SandboxConfig.Hostname,
		"SandboxNamespace": in.SandboxConfig.Metadata.Namespace,
	}).Debug("CreateContainer")
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

// getSandboxRootDir returns the root directory for managing sandbox files,
// e.g. hosts files.
func (s *RuntimeServer) getSandboxRootDir(id string) string {
	return filepath.Join(rootDir, sandboxesDir, id)
}
