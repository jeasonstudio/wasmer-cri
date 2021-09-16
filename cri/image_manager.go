package cri

import runtimeapi "github.com/jeasonstudio/wasmer-cri/cri/runtime/v1alpha2"

// IImageManager interface should be implemented by a container image
// manager.
// The methods should be thread-safe.
type IImageManager interface {
	// ListImages lists the existing images.
	ListImages(filter *runtimeapi.ImageFilter) ([]*runtimeapi.Image, error)
	// ImageStatus returns the status of the image.
	ImageStatus(image *runtimeapi.ImageSpec) (*runtimeapi.Image, error)
	// PullImage pulls an image with the authentication config.
	PullImage(image *runtimeapi.ImageSpec, auth *runtimeapi.AuthConfig, podSandboxConfig *runtimeapi.PodSandboxConfig) (string, error)
	// RemoveImage removes the image.
	RemoveImage(image *runtimeapi.ImageSpec) error
	// ImageFsInfo returns information of the filesystem that is used to store images.
	ImageFsInfo() ([]*runtimeapi.FilesystemUsage, error)
}

type ImageManager struct{}

// NewImageManager create ImageManager instance
func NewImageManager() (*ImageManager, error) {
	mgr := &ImageManager{}
	return mgr, nil
}

// ListImages lists the existing images.
func (im *ImageManager) ListImages(filter *runtimeapi.ImageFilter) ([]*runtimeapi.Image, error) {
	return nil, nil
}

// ImageStatus returns the status of the image.
func (im *ImageManager) ImageStatus(image *runtimeapi.ImageSpec) (*runtimeapi.Image, error) {
	return nil, nil
}

// PullImage pulls an image with the authentication config.
func (im *ImageManager) PullImage(image *runtimeapi.ImageSpec, auth *runtimeapi.AuthConfig, podSandboxConfig *runtimeapi.PodSandboxConfig) (string, error) {
	return "", nil
}

// RemoveImage removes the image.
func (im *ImageManager) RemoveImage(image *runtimeapi.ImageSpec) error {
	return nil
}

// ImageFsInfo returns information of the filesystem that is used to store images.
func (im *ImageManager) ImageFsInfo() ([]*runtimeapi.FilesystemUsage, error) {
	return nil, nil
}
