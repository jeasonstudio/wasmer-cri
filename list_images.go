package wasmercri

import (
	"context"

	log "github.com/sirupsen/logrus"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// ListImages list images
func (s *ImageServer) ListImages(ctx context.Context, in *pb.ListImagesRequest) (*pb.ListImagesResponse, error) {
	logger := log.WithContext(ctx)
	logger.WithFields(log.Fields{
		"image":       in.Filter.Image.Image,
		"annotations": in.Filter.Image.Annotations,
	}).Debug("ListImages")

	id := "sha256:E58FCF7418D4390DEC8E8FB69D88C06EC07039D651FEDD3AA72AF9972E7D046B"

	return &pb.ListImagesResponse{Images: []*pb.Image{{Id: id, RepoTags: []string{"ghcr.io/jeasonstudio/example.wasm:latest"}, Size_: 10000, Username: "jeason", RepoDigests: []string{}, Spec: &pb.ImageSpec{
		Image:       "ghcr.io/jeasonstudio/example.wasm:latest",
		Annotations: map[string]string{},
	}}}}, nil
}
