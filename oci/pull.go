package oci

import (
	"context"
	"io/ioutil"
	"log"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
)

// Pull pulls a Wasm module from an OCI registry given a reference
func Pull(ref, filename string, config *RegistryConfig) (*v1.Descriptor, error) {
	ctx := context.Background()
	store := content.NewMemory()

	registry, err := NewRegistry(config)
	if err != nil {
		log.Fatalf("Fail to pull image, %s", err)
		return nil, err
	}
	var descriptor v1.Descriptor
	_, err = oras.Copy(ctx, registry, ref, store, "", oras.WithAllowedMediaTypes([]string{ContentLayerMediaType, ConfigMediaType}), oras.WithPullEmptyNameAllowed(), oras.WithLayerDescriptors(func(descriptors []v1.Descriptor) {
		descriptor = descriptors[0]
		_, contents, _ := store.Get(descriptor)
		ioutil.WriteFile(filename, contents, 0755)
	}))
	if err != nil {
		log.Fatalf("Fail to pull image, %s", err)
		return nil, err
	}

	return &descriptor, nil
}
