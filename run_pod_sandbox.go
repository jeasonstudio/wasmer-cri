package wasmercri

import (
	"context"
	"time"

	"github.com/jeasonstudio/wasmer-cri/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

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
	name := makeSandboxName(metadata)
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
