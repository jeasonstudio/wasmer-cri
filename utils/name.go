package utils

import (
	"fmt"
	"strings"

	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

const (
	// Delimiter used to construct container/sandbox names.
	nameDelimiter = "_"
)

// MakeSandboxName generates sandbox name from sandbox metadata. The name
// generated is unique as long as sandbox metadata is unique.
func MakeSandboxName(s *pb.PodSandboxMetadata) string {
	return strings.Join([]string{
		s.Name,                       // 0
		s.Namespace,                  // 1
		s.Uid,                        // 2
		fmt.Sprintf("%d", s.Attempt), // 3
	}, nameDelimiter)
}
