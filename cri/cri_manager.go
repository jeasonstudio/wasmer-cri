package cri

import (
	"fmt"

	runtimeapi "github.com/jeasonstudio/wasmer-cri/cri/runtime/v1alpha2"
)

type CRIManager struct {
	ContainerManager ContainerManager

	// SandboxBaseDir is the directory used to store sandbox files like /etc/hosts, /etc/resolv.conf, etc.
	SandboxBaseDir string
	// SandboxImage is the image used by sandbox container.
	SandboxImage string
}

// Version returns the runtime name, runtime version and runtime API version.
func (c *CRIManager) Version(apiVersion string) (*runtimeapi.VersionResponse, error) {
	// TODO: constants should be refactor
	return &runtimeapi.VersionResponse{
		Version:           "1.0.0",
		RuntimeName:       "Wasmer",
		RuntimeVersion:    "1.0.0",
		RuntimeApiVersion: "1.0.0",
	}, nil
}

// RunPodSandbox creates and starts a pod-level sandbox. Runtimes should ensure the sandbox is in ready state.
func (c *CRIManager) RunPodSandbox(config *runtimeapi.PodSandboxConfig, runtimeHandler string) (_ string, err error) {
	containerId := "container-pod-id" // TODO: needs to be generated

	// Step 1: Prepare image for the sandbox.
	// Make sure the sandbox image exists.
	err = c.ensureSandboxImageExists(c.SandboxImage)
	if err != nil {
		return "", err
	}

	// Step 2: Setup networking for the sandbox.
	// TODO

	// Step 3: Create the sandbox container.
	// TODO: containerConfig := &runtimeapi.ContainerConfig{}
	_, err = c.ContainerManager.CreateContainer(containerId, nil, config)
	if err != nil {
		return "", fmt.Errorf("failed to create a sandbox for pod %q: %v", config.Metadata.Name, err)
	}

	// If running sandbox failed, clean up the container.
	defer func() {
		if err != nil {
			c.ContainerManager.RemoveContainer(containerId)
		}
	}()

	// Step 4: Start the sandbox container.
	err = c.ContainerManager.StartContainer(containerId)
	if err != nil {
		return "", fmt.Errorf("failed to start sandbox container for pod %q: %v", containerId, err)
	}

	return "", nil
}
