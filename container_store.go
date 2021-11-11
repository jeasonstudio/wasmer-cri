package wasmercri

import (
	"errors"
	"sync"

	"github.com/jeasonstudio/wasmer-cri/truncindex"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// ContainerStore store for pod sandbox
type ContainerStore struct {
	lock      sync.RWMutex
	sandboxes map[string]pb.Container
	idIndex   *truncindex.TruncIndex
}

// NewContainerStore create pod-sandbox store
func NewContainerStore() *ContainerStore {
	return &ContainerStore{
		sandboxes: make(map[string]pb.Container),
		idIndex:   truncindex.NewTruncIndex([]string{}),
	}
}

// Add a sandbox into the store.
func (s *ContainerStore) Add(sb *pb.Container) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	if _, ok := s.sandboxes[sb.Id]; ok {
		return errors.New("pod already exists")
	}
	if err := s.idIndex.Add(sb.Id); err != nil {
		return err
	}
	if _, ok := s.sandboxes[sb.Id]; ok {
		return errors.New("pod already exists")
	}
	s.sandboxes[sb.Id] = *sb
	return nil
}

// Get returns the sandbox with specified id.
func (s *ContainerStore) Get(id string) (pb.Container, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	id, err := s.idIndex.Get(id)
	if err != nil {
		if err == truncindex.ErrNotExist {
			err = errors.New("pod not found")
		}
		return pb.Container{}, err
	}

	if sb, ok := s.sandboxes[id]; ok {
		return sb, nil
	}
	return pb.Container{}, errors.New("pod not found")
}

// List lists all sandboxes.
func (s *ContainerStore) List() []pb.Container {
	s.lock.RLock()
	defer s.lock.RUnlock()

	var sandboxes []pb.Container

	for _, sb := range s.sandboxes {
		sandboxes = append(sandboxes, sb)
	}
	return sandboxes
}

// Delete deletes the sandbox with specified id.
func (s *ContainerStore) Delete(id string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	id, err := s.idIndex.Get(id)
	if err != nil {
		// Note: The idIndex.Delete and delete doesn't handle truncated index.
		// So we need to return if there are error.
		return
	}
	s.idIndex.Delete(id) // nolint: errcheck
	delete(s.sandboxes, id)
}
