package wasmercri

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/distribution/distribution/v3/reference"
	imagedigest "github.com/opencontainers/go-digest"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

const (
	// errorStartReason is the exit reason when fails to start container.
	errorStartReason = "StartError"
	// errorStartExitCode is the exit code when fails to start container.
	// 128 is the same with Docker's behavior.
	// TODO(windows): Figure out what should be used for windows.
	errorStartExitCode = 128
	// completeExitReason is the exit reason when container exits with code 0.
	completeExitReason = "Completed"
	// errorExitReason is the exit reason when container exits with code non-zero.
	errorExitReason = "Error"
	// oomExitReason is the exit reason when process in container is oom killed.
	oomExitReason = "OOMKilled"

	// sandboxesDir contains all sandbox root. A sandbox root is the running
	// directory of the sandbox, all files created for the sandbox will be
	// placed under this directory.
	sandboxesDir = "sandboxes"
	// containersDir contains all container root.
	containersDir = "containers"
	// Delimiter used to construct container/sandbox names.
	nameDelimiter = "_"

	// criWasmerdPrefix is common prefix for cri-wasmerd
	criWasmerdPrefix = "io.cri-wasmerd"
	// containerKindLabel is a label key indicating container is sandbox container or application container
	containerKindLabel = criWasmerdPrefix + ".kind"
	// containerKindSandbox is a label value indicating container is sandbox container
	containerKindSandbox = "sandbox"
	// containerKindContainer is a label value indicating container is application container
	containerKindContainer = "container"
	// imageLabelKey is the label key indicating the image is managed by cri plugin.
	imageLabelKey = criWasmerdPrefix + ".image"
	// imageLabelValue is the label value indicating the image is managed by cri plugin.
	imageLabelValue = "managed"
	// sandboxMetadataExtension is an extension name that identify metadata of sandbox in CreateContainerRequest
	sandboxMetadataExtension = criWasmerdPrefix + ".sandbox.metadata"
	// containerMetadataExtension is an extension name that identify metadata of container in CreateContainerRequest
	containerMetadataExtension = criWasmerdPrefix + ".container.metadata"

	// defaultIfName is the default network interface for the pods
	defaultIfName = "eth0"

	// runtimeRunhcsV1 is the runtime type for runhcs.
	// runtimeRunhcsV1 = "io.wasmerd.runhcs.v1"
)

const (
	rootDir = ".wasmerd"
)

// makeSandboxName generates sandbox name from sandbox metadata. The name
// generated is unique as long as sandbox metadata is unique.
func makeSandboxName(s *pb.PodSandboxMetadata) string {
	return strings.Join([]string{
		s.Name,                       // 0
		s.Namespace,                  // 1
		s.Uid,                        // 2
		fmt.Sprintf("%d", s.Attempt), // 3
	}, nameDelimiter)
}

// makeContainerName generates container name from sandbox and container metadata.
// The name generated is unique as long as the sandbox container combination is
// unique.
func makeContainerName(c *pb.ContainerMetadata, s *pb.PodSandboxMetadata) string {
	return strings.Join([]string{
		c.Name,                       // 0: container name
		s.Name,                       // 1: pod name
		s.Namespace,                  // 2: pod namespace
		s.Uid,                        // 3: pod uid
		fmt.Sprintf("%d", c.Attempt), // 4: attempt number of creating the container
	}, nameDelimiter)
}

// getSandboxRootDir returns the root directory for managing sandbox files,
// e.g. hosts files.
func (s *RuntimeServer) getSandboxRootDir(id string) string {
	return filepath.Join(rootDir, sandboxesDir, id)
}

// getRepoDigestAngTag returns image repoDigest and repoTag of the named image reference.
func getRepoDigestAndTag(namedRef reference.Named, digest imagedigest.Digest, schema1 bool) (string, string) {
	var repoTag, repoDigest string
	if _, ok := namedRef.(reference.NamedTagged); ok {
		repoTag = namedRef.String()
	}
	if _, ok := namedRef.(reference.Canonical); ok {
		repoDigest = namedRef.String()
	} else if !schema1 {
		// digest is not actual repo digest for schema1 image.
		repoDigest = namedRef.Name() + "@" + digest.String()
	}
	return repoDigest, repoTag
}
