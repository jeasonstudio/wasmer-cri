package wasmercri

import (
	"context"

	log "github.com/sirupsen/logrus"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

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
