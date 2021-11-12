package wasmercri

import (
	"context"

	"github.com/distribution/distribution/v3/reference"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// PullImage pull image
func (s *ImageServer) PullImage(ctx context.Context, in *pb.PullImageRequest) (*pb.PullImageResponse, error) {
	logger := log.WithContext(ctx)
	logger.WithFields(log.Fields{
		"image":         in.Image.Image,
		"auth":          in.Auth,
		"sandboxConfig": in.SandboxConfig,
	}).Debug("ListImages")

	imageRef := in.GetImage().GetImage()
	logger.WithField("imageRef", imageRef).Debug("generate image ref")

	namedRef, err := reference.ParseDockerRef(imageRef)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse image reference %q", imageRef)
	}
	logger.WithField("namedRef", namedRef).Debug("generate named ref")

	ref := namedRef.String()
	if ref != imageRef {
		logger.Debugf("pull image using normalized image ref: %q", ref)
	}
	return &pb.PullImageResponse{
		ImageRef: "my-image-ref",
	}, nil
}
