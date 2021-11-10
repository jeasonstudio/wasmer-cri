package oci

import "oras.land/oras-go/pkg/content"

// RegistryConfig oci registry config
type RegistryConfig struct {
	content.RegistryOptions
	// TODO: may append some configs
}

// NewRegistry create registry instance
func NewRegistry(config *RegistryConfig) (*content.Registry, error) {
	options := &content.RegistryOptions{}
	if config != nil {
		options.Configs = config.Configs
		options.Username = config.Username
		options.Password = config.Password
		options.Insecure = config.Insecure
		options.PlainHTTP = config.PlainHTTP
	}
	registry, err := content.NewRegistry(*options)
	return registry, err
}
