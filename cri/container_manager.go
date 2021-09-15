package cri

import (
	"time"

	runtimeapi "github.com/jeasonstudio/wasmer-cri/cri/runtime/v1alpha2"
)

// IContainerManager contains methods to manipulate containers managed by a
// container runtime. The methods are thread-safe.
type IContainerManager interface {
	// CreateContainer creates a new container in specified PodSandbox.
	CreateContainer(podSandboxID string, config *runtimeapi.ContainerConfig, sandboxConfig *runtimeapi.PodSandboxConfig) (string, error)
	// StartContainer starts the container.
	StartContainer(containerID string) error
	// StopContainer stops a running container with a grace period (i.e., timeout).
	StopContainer(containerID string, timeout int64) error
	// RemoveContainer removes the container.
	RemoveContainer(containerID string) error
	// ListContainers lists all containers by filters.
	ListContainers(filter *runtimeapi.ContainerFilter) ([]*runtimeapi.Container, error)
	// ContainerStatus returns the status of the container.
	ContainerStatus(containerID string) (*runtimeapi.ContainerStatus, error)
	// UpdateContainerResources updates the cgroup resources for the container.
	UpdateContainerResources(containerID string, resources *runtimeapi.LinuxContainerResources) error
	// ExecSync executes a command in the container, and returns the stdout output.
	// If command exits with a non-zero exit code, an error is returned.
	ExecSync(containerID string, cmd []string, timeout time.Duration) (stdout []byte, stderr []byte, err error)
	// Exec prepares a streaming endpoint to execute a command in the container, and returns the address.
	Exec(*runtimeapi.ExecRequest) (*runtimeapi.ExecResponse, error)
	// Attach prepares a streaming endpoint to attach to a running container, and returns the address.
	Attach(req *runtimeapi.AttachRequest) (*runtimeapi.AttachResponse, error)
	// ReopenContainerLog asks runtime to reopen the stdout/stderr log file
	// for the container. If it returns error, new container log file MUST NOT
	// be created.
	ReopenContainerLog(ContainerID string) error
}

// ContainerManager is the implement of interface IContainerManager.
type ContainerManager struct{}

// NewContainerManager create ContainerManager instance
func NewContainerManager() (*ContainerManager, error) {
	mgr := &ContainerManager{}
	return mgr, nil
}

// CreateContainer creates a new container in specified PodSandbox.
func (cm *ContainerManager) CreateContainer(podSandboxID string, config *runtimeapi.ContainerConfig, sandboxConfig *runtimeapi.PodSandboxConfig) (string, error) {
	return "", nil
}

// StartContainer starts the container.
func (cm *ContainerManager) StartContainer(containerID string) error {
	return nil
}

// StopContainer stops a running container with a grace period (i.e., timeout).
func (cm *ContainerManager) StopContainer(containerID string, timeout int64) error {
	return nil
}

// RemoveContainer removes the container.
func (cm *ContainerManager) RemoveContainer(containerID string) error {
	return nil
}

// ListContainers lists all containers by filters.
func (cm *ContainerManager) ListContainers(filter *runtimeapi.ContainerFilter) ([]*runtimeapi.Container, error) {
	return nil, nil
}

// ContainerStatus returns the status of the container.
func (cm *ContainerManager) ContainerStatus(containerID string) (*runtimeapi.ContainerStatus, error) {
	return nil, nil
}

// Exec prepares a streaming endpoint to execute a command in the container, and returns the address.
func (cm *ContainerManager) Exec(*runtimeapi.ExecRequest) (*runtimeapi.ExecResponse, error) {
	return nil, nil
}
