package config

// TODO

const (
	// K8sNamespace is the namespace we use to connect containerd when CRI is enabled.
	K8sNamespace = "k8s.io"
)

// Config defines the CRI configuration.
type Config struct {
	// Listen is the listening address which servers CRI.
	Listen string `json:"listen,omitempty"`
}
