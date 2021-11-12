package oci

import v1 "github.com/opencontainers/image-spec/specs-go/v1"

// Client OCI client
type Client struct {
	options *options
}

// NewClient create oci client
func NewClient() *Client {
	return &Client{
		options: &options{
			configs: []string{},
		},
	}
}

// OptPullPush options for pull/push
type OptPullPush func(o *options) error

type options struct {
	// registry
	configs             []string
	username, password  string
	insecure, plainHTTP bool
	// handleManifestDescriptor func([]byte)
	handleManifestDescriptor func(v1.Descriptor)
	handleLayerDescriptor    func(v1.Descriptor)
	// saveLayers   func([]ocispec.Descriptor)
	// validateName func(desc ocispec.Descriptor) error
}

// WithNetworkConfig sets the allowed media types
func WithNetworkConfig(insecure, plainHTTP bool) OptPullPush {
	return func(o *options) error {
		o.insecure = insecure
		o.plainHTTP = plainHTTP
		return nil
	}
}

// WithRegistryConfig sets the allowed media types
func WithRegistryConfig(insecure, plainHTTP bool) OptPullPush {
	return func(o *options) error {
		o.insecure = insecure
		o.plainHTTP = plainHTTP
		return nil
	}
}

// WithHandleManifestDesc sets the allowed media types
func WithHandleManifestDesc(save func(manifest v1.Descriptor)) OptPullPush {
	return func(o *options) error {
		o.handleManifestDescriptor = save
		return nil
	}
}

// WithHandleLayerDesc sets the allowed media types
func WithHandleLayerDesc(save func(layer v1.Descriptor)) OptPullPush {
	return func(o *options) error {
		o.handleLayerDescriptor = save
		return nil
	}
}
