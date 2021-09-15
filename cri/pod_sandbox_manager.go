package cri

import (
	"fmt"

	runtimeapi "github.com/jeasonstudio/wasmer-cri/cri/runtime/v1alpha2"
)

var (
	// Default timeout for stopping container.
	defaultStopTimeout = int64(10)
)

// PodSandboxManager contains methods for operating on PodSandboxes. The methods
// are thread-safe.
type IPodSandboxManager interface {
	// RunPodSandbox creates and starts a pod-level sandbox. Runtimes should ensure
	// the sandbox is in ready state.
	RunPodSandbox(config *runtimeapi.PodSandboxConfig, runtimeHandler string) (string, error)
	// StopPodSandbox stops the sandbox. If there are any running containers in the
	// sandbox, they should be force terminated.
	StopPodSandbox(podSandboxID string) error
	// RemovePodSandbox removes the sandbox. If there are running containers in the
	// sandbox, they should be forcibly removed.
	RemovePodSandbox(podSandboxID string) error
	// PodSandboxStatus returns the Status of the PodSandbox.
	PodSandboxStatus(podSandboxID string) (*runtimeapi.PodSandboxStatus, error)
	// ListPodSandbox returns a list of Sandbox.
	ListPodSandbox(filter *runtimeapi.PodSandboxFilter) ([]*runtimeapi.PodSandbox, error)
	// PortForward prepares a streaming endpoint to forward ports from a PodSandbox, and returns the address.
	PortForward(*runtimeapi.PortForwardRequest) (*runtimeapi.PortForwardResponse, error)
}

type PodSandboxManager struct {
	ContainerManager ContainerManager

	// SandboxBaseDir is the directory used to store sandbox files like /etc/hosts, /etc/resolv.conf, etc.
	SandboxBaseDir string
	// SandboxImage is the image used by sandbox container.
	SandboxImage string
}

// ensureSandboxImageExists pulls the image when it's not present.
func (psm *PodSandboxManager) ensureSandboxImageExists(imageRef string) error {
	// TODO: needs image service
	// _, _, _, err := c.ImageMgr.CheckReference(ctx, imageRef)
	// if err == nil {
	// 	return nil
	// }
	// if errtypes.IsNotfound(err) {
	// 	err = c.ImageMgr.PullImage(ctx, imageRef, nil, bytes.NewBuffer([]byte{}))
	// 	if err != nil {
	// 		return fmt.Errorf("failed to pull sandbox image %q: %v", imageRef, err)
	// 	}
	// 	return nil
	// }
	// return fmt.Errorf("failed to check sandbox image %q: %v", imageRef, err)
	return nil
}

// Version returns the runtime name, runtime version and runtime API version.
func (psm *PodSandboxManager) Version(apiVersion string) (*runtimeapi.VersionResponse, error) {
	// TODO: constants should be refactor
	return &runtimeapi.VersionResponse{
		Version:           "1.0.0",
		RuntimeName:       "Wasmer",
		RuntimeVersion:    "1.0.0",
		RuntimeApiVersion: "1.0.0",
	}, nil
}

// RunPodSandbox creates and starts a pod-level sandbox. Runtimes should ensure the sandbox is in ready state.
func (psm *PodSandboxManager) RunPodSandbox(config *runtimeapi.PodSandboxConfig, runtimeHandler string) (_ string, err error) {
	containerId := "container-pod-id" // TODO: needs to be generated

	// Step 1: Prepare image for the sandbox.
	// Make sure the sandbox image exists.
	err = psm.ensureSandboxImageExists(psm.SandboxImage)
	if err != nil {
		return "", err
	}

	// Step 2: Setup networking for the sandbox.
	// TODO

	// Step 3: Create the sandbox container.
	// TODO: containerConfig := &runtimeapi.ContainerConfig{}
	_, err = psm.ContainerManager.CreateContainer(containerId, nil, config)
	if err != nil {
		return "", fmt.Errorf("failed to create a sandbox for pod %q: %v", config.Metadata.Name, err)
	}

	// If running sandbox failed, clean up the container.
	defer func() {
		if err != nil {
			psm.ContainerManager.RemoveContainer(containerId)
		}
	}()

	// Step 4: Start the sandbox container.
	err = psm.ContainerManager.StartContainer(containerId)
	if err != nil {
		return "", fmt.Errorf("failed to start sandbox container for pod %q: %v", containerId, err)
	}

	return "", nil
}

// StopPodSandbox stops the sandbox. If there are any running containers in the sandbox, they should be force terminated.
func (psm *PodSandboxManager) StopPodSandbox(podSandboxID string) (err error) {
	var containers []*runtimeapi.Container

	containers, err = psm.ContainerManager.ListContainers(nil) // nil means all containers under this pod sandbox
	if err != nil {
		return fmt.Errorf("failed to get the containers belong to sandbox %q: %v", podSandboxID, err)
	}

	// Stop all containers in the sandbox.
	for _, container := range containers {
		err = psm.ContainerManager.StopContainer(container.Id, defaultStopTimeout)
		if err != nil {
			return fmt.Errorf("failed to stop container %q of sandbox %q: %v", container.Id, podSandboxID, err)
		}
		fmt.Printf("Success to stop container %q of sandbox %q", container.Id, podSandboxID)
	}

	// TODO: stop the network settings

	return nil
}

// RemovePodSandbox removes the sandbox. If there are running containers in the sandbox, they should be forcibly removed.
func (psm *PodSandboxManager) RemovePodSandbox(podSandboxID string) error {
	return nil
}

// PodSandboxStatus returns the Status of the PodSandbox.
func (psm *PodSandboxManager) PodSandboxStatus(podSandboxID string) (*runtimeapi.PodSandboxStatus, error) {
	return nil, nil
}

// ListPodSandbox returns a list of Sandbox.
func (psm *PodSandboxManager) ListPodSandbox(filter *runtimeapi.PodSandboxFilter) ([]*runtimeapi.PodSandbox, error) {
	return nil, nil
}
