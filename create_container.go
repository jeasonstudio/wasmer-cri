package wasmercri

import (
	"context"

	log "github.com/sirupsen/logrus"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

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
