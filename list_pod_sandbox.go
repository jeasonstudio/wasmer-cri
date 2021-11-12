package wasmercri

import (
	"context"

	log "github.com/sirupsen/logrus"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

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
