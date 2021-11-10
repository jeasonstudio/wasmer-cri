package oci

import (
	"context"
	"io/ioutil"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
)

// Pull pulls a Wasm module from an OCI registry given a reference
func Pull(ref string, config *RegistryConfig) (*v1.Descriptor, []byte, error) {
	log.WithFields(log.Fields{
		"ref":            ref,
		"registryConfig": config,
	}).Debug("Pull pure image")

	ctx := context.Background()
	store := content.NewMemory()

	registry, err := NewRegistry(config)
	if err != nil {
		log.WithError(err).Fatal("Failed to create register when pull image")
		return nil, nil, err
	}
	var descriptor v1.Descriptor
	var contents []byte
	_, err = oras.Copy(ctx, registry, ref, store, "", oras.WithAllowedMediaTypes([]string{ContentLayerMediaType, ConfigMediaType}), oras.WithPullEmptyNameAllowed(), oras.WithLayerDescriptors(func(descriptors []v1.Descriptor) {
		descriptor = descriptors[0] // Use the top layer
		_, contents, _ = store.Get(descriptor)
	}), oras.WithNameValidation(func(desc v1.Descriptor) error {
		// TODO: validata name should be ends with ".wasm"
		return nil
	}))

	if err != nil {
		log.WithError(err).Fatal("Failed to pull image")
		return nil, nil, err
	}

	log.WithFields(log.Fields{
		"mediaType":  descriptor.MediaType,
		"descriptor": descriptor,
	}).Debug("ImageInfo")

	return &descriptor, contents, nil
}

// PullToFile Pull and write to file system
func PullToFile(ref, filename string, config *RegistryConfig) (*v1.Descriptor, error) {
	log.WithFields(log.Fields{
		"ref":            ref,
		"registryConfig": config,
	}).Debug("Pull image to file")

	descriptor, contents, err := Pull(ref, config)
	if err != nil {
		log.WithError(err).Fatal("Failed to pull image")
		return nil, err
	}

	err = ioutil.WriteFile(filename, contents, 0755)
	if err != nil {
		log.WithError(err).Fatal("Failed to write image into file-system")
		return nil, err
	}
	return descriptor, nil
}
