package oci

import (
	"context"
	"errors"
	"strings"

	"github.com/distribution/distribution/v3/reference"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
)

// Pull pulls a Wasm module from an OCI registry given a reference
func (c *Client) Pull(ctx context.Context, ref string, middlewares ...OptPullPush) (*Image, error) {
	logger := log.WithContext(ctx)
	logger.WithField("ref", ref).Debug("[oci] pull image")

	// Verify ref for name
	refName, err := reference.ParseDockerRef(ref)
	if err != nil {
		logger.WithError(err).Fatal("[oci] failed to parse ref")
		return nil, err
	}
	logger.Debugln(refName)

	// Create store
	store := content.NewMemory()

	// Apply all middlewares
	for _, middleware := range middlewares {
		if err := middleware(c.options); err != nil {
			return nil, err
		}
	}

	var descriptor v1.Descriptor
	var contents []byte

	// Pull options
	pullOpts := []oras.CopyOpt{
		oras.WithAllowedMediaTypes([]string{MediaTypeWasmContentLayer, MediaTypeWasmConfig}),
		oras.WithPullEmptyNameAllowed(),
		oras.WithLayerDescriptors(func(descriptors []v1.Descriptor) {
			descriptor = descriptors[0] // Use the top layer
			_, contents, _ = store.Get(descriptor)
		}),
		oras.WithNameValidation(func(desc v1.Descriptor) error {
			if strings.HasSuffix(desc.Annotations[v1.AnnotationTitle], ".wasm") {
				return nil
			}
			return errors.New("Invalidate content name, should be ends with .wasm")
		}),
	}

	// registry
	registryOptions := &content.RegistryOptions{
		Username:  c.options.username,
		Password:  c.options.password,
		Configs:   c.options.configs,
		Insecure:  c.options.insecure,
		PlainHTTP: c.options.plainHTTP,
	}
	registry, err := content.NewRegistry(*registryOptions)
	if err != nil {
		logger.WithError(err).Fatal("[oci] failed to create registry")
		return nil, err
	}

	rootManifest, err := oras.Copy(ctx, registry, ref, store, "", pullOpts...)
	if err != nil {
		logger.WithError(err).Fatal("[oci] failed to pull image")
		return nil, err
	}

	image := &Image{
		Name:        refName.Name(),
		Labels:      map[string]string{},
		Manifest:    rootManifest,
		Webassembly: descriptor,
		Content:     contents,
	}

	return image, nil
}
