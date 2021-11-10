package wasmercri

import (
	"errors"
	"sync"
	"time"

	"github.com/jeasonstudio/wasmer-cri/truncindex"
	pb "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

// PodSandboxStore store for pod sandbox
type PodSandboxStore struct {
	lock      sync.RWMutex
	sandboxes map[string]pb.PodSandbox
	idIndex   *truncindex.TruncIndex
}

// NewSandbox create pod sandbox
func NewSandbox(id string, metadata *pb.PodSandboxMetadata, state pb.PodSandboxState) *pb.PodSandbox {
	pod := &pb.PodSandbox{
		Id:             id,
		Metadata:       metadata,
		State:          state,
		CreatedAt:      time.Now().UnixNano(),
		Labels:         map[string]string{},
		Annotations:    map[string]string{},
		RuntimeHandler: "",
	}
	return pod
}

// NewPodSandboxStore create pod-sandbox store
func NewPodSandboxStore() *PodSandboxStore {
	return &PodSandboxStore{
		sandboxes: make(map[string]pb.PodSandbox),
		idIndex:   truncindex.NewTruncIndex([]string{}),
	}
}

// Add a sandbox into the store.
func (s *PodSandboxStore) Add(sb pb.PodSandbox) error {
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
	s.sandboxes[sb.Id] = sb
	return nil
}

// Get returns the sandbox with specified id.
func (s *PodSandboxStore) Get(id string) (*pb.PodSandbox, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	id, err := s.idIndex.Get(id)
	if err != nil {
		if err == truncindex.ErrNotExist {
			err = errors.New("pod not found")
		}
		return nil, err
	}

	if sb, ok := s.sandboxes[id]; ok {
		return &sb, nil
	}
	return nil, errors.New("pod not found")
}

// List lists all sandboxes.
func (s *PodSandboxStore) List() []*pb.PodSandbox {
	s.lock.RLock()
	defer s.lock.RUnlock()

	var sandboxes []*pb.PodSandbox

	for _, sb := range s.sandboxes {
		sandboxes = append(sandboxes, &sb)
	}
	return sandboxes
}

// Delete deletes the sandbox with specified id.
func (s *PodSandboxStore) Delete(id string) {
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
