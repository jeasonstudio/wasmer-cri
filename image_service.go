package wasmercri

import (
	"context"

	"time"

	log "github.com/sirupsen/logrus"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// ImageServer is used to implement images
type ImageServer struct {
	pb.UnimplementedImageServiceServer
}

// NewImageServer create image server
func NewImageServer() (*ImageServer, error) {
	log.Trace("NewImageServer create image server")
	return &ImageServer{}, nil
}

// ListImages list images
func (s *ImageServer) ListImages(ctx context.Context, in *pb.ListImagesRequest) (*pb.ListImagesResponse, error) {
	log.WithFields(log.Fields{
		"image":       in.Filter.Image.Image,
		"annotations": in.Filter.Image.Annotations,
	}).Debug("ListImages")

	id := "sha256:E58FCF7418D4390DEC8E8FB69D88C06EC07039D651FEDD3AA72AF9972E7D046B"

	return &pb.ListImagesResponse{Images: []*pb.Image{{Id: id, RepoTags: []string{"ghcr.io/jeasonstudio/example.wasm:latest"}, Size_: 10000, Username: "jeason", RepoDigests: []string{}, Spec: &pb.ImageSpec{
		Image:       "ghcr.io/jeasonstudio/example.wasm:latest",
		Annotations: map[string]string{},
	}}}}, nil
}

// ImageStatus show status of image
func (s *ImageServer) ImageStatus(ctx context.Context, in *pb.ImageStatusRequest) (*pb.ImageStatusResponse, error) {
	log.Printf("ImageStatus Received: %v", in.Image)
	myImg := &pb.Image{Id: "id2", RepoTags: []string{}, Size_: 10000, Username: "foo/bar", RepoDigests: []string{}, Spec: &pb.ImageSpec{
		Image:       "wasm/hello-world",
		Annotations: map[string]string{},
	}}
	return &pb.ImageStatusResponse{
		Image: myImg,
		Info:  map[string]string{},
	}, nil
}

// PullImage pull image
func (s *ImageServer) PullImage(ctx context.Context, in *pb.PullImageRequest) (*pb.PullImageResponse, error) {
	log.Printf("PullImage Received: %v", in.Image)
	return &pb.PullImageResponse{
		ImageRef: "my-image-ref",
	}, nil
}

// RemoveImage remove image
func (s *ImageServer) RemoveImage(ctx context.Context, in *pb.RemoveImageRequest) (*pb.RemoveImageResponse, error) {
	log.Printf("RemoveImage Received: %v", in.Image)
	return &pb.RemoveImageResponse{}, nil
}

// ImageFsInfo get image file-system info
func (s *ImageServer) ImageFsInfo(ctx context.Context, in *pb.ImageFsInfoRequest) (*pb.ImageFsInfoResponse, error) {
	log.Printf("ImageFsInfo Received: %v", in)
	fs := &pb.FilesystemUsage{
		Timestamp: time.Now().UnixNano(),
		FsId: &pb.FilesystemIdentifier{
			Mountpoint: "$HOME/.wasmer",
		},
		UsedBytes:  &pb.UInt64Value{Value: 0},
		InodesUsed: &pb.UInt64Value{Value: 0},
	}
	return &pb.ImageFsInfoResponse{
		ImageFilesystems: []*pb.FilesystemUsage{fs},
	}, nil
}
