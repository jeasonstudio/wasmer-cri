package wasmercri

import (
	"sync"

	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// ImageStore store for images
type ImageStore struct {
	lock sync.RWMutex
	// refCache is a containerd image reference to image id cache.
	cache map[string]string
	// store is the internal image store indexed by image id.
	// store *store
}

// NewImageStore creates an image store.
func NewImageStore() *ImageStore {
	return &ImageStore{
		cache: make(map[string]string),
	}
}

// Get get image by id
func (s *ImageStore) Get(id string) (*pb.Image, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return nil, nil
}
