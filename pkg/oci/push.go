package oci

import (
	"context"
	"errors"
	"strings"
	"time"

	"github.com/distribution/distribution/v3/reference"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
)

// Push pushes a WASM module to an OCI registry
func (c *Client) Push(ctx context.Context, ref, filename string, moduleContent []byte, middlewares ...OptPullPush) (*Image, error) {
	logger := log.WithContext(ctx)
	logger.WithFields(log.Fields{
		"ref":      ref,
		"filename": filename,
	}).Debug("[oci] push params")

	// Verify ref for name
	refName, err := reference.ParseDockerRef(ref)
	if err != nil {
		logger.WithError(err).Fatal("[oci] failed to parse ref")
		return nil, err
	}

	// Create store
	store := content.NewMemory()

	// Apply all middlewares
	for _, middleware := range middlewares {
		if err := middleware(c.options); err != nil {
			return nil, err
		}
	}

	// Push options
	pushOpts := []oras.CopyOpt{
		oras.WithAllowedMediaTypes([]string{MediaTypeWasmContentLayer, MediaTypeWasmConfig}),
		oras.WithLayerDescriptors(func(layers []v1.Descriptor) {
			for _, layer := range layers {
				logger.WithField("digest", layer.Digest).Info("Pushing image layer...")
			}
		}),
		oras.WithNameValidation(func(desc v1.Descriptor) error {
			if strings.HasSuffix(desc.Annotations[v1.AnnotationTitle], ".wasm") {
				return nil
			}
			return errors.New("Invalidate content name, should be ends with .wasm")
		}),
	}

	configAnnotations := map[string]string{
		// TODO: add more annotation
		AnnotationWasmTitle:  filename,
		v1.AnnotationCreated: time.Now().Format(time.RFC3339),
	}
	manifestAnnotations := map[string]string{
		// TODO: add more annotation
		AnnotationWasmTitle:  filename,
		v1.AnnotationCreated: time.Now().Format(time.RFC3339),
	}

	coreWasmContentLayerDesc, err := store.Add("", MediaTypeWasmContentLayer, moduleContent)
	if err != nil {
		logger.WithError(err).Fatal("[oci] failed to read Webassembly file")
		return nil, err
	}
	coreWasmContentLayerDesc.Annotations = map[string]string{
		AnnotationWasmTitle: filename,
	}

	// Webassembly content only has one layer
	layers := []v1.Descriptor{
		coreWasmContentLayerDesc,
	}

	// Generate OCI config
	config, configDescriptor, err := content.GenerateConfig(configAnnotations)
	if err != nil {
		logger.WithError(err).Fatal("[oci] failed to generate config")
		return nil, err
	}
	configDescriptor.MediaType = MediaTypeWasmConfig // change to wasm media type in config
	store.Set(configDescriptor, config)

	// Generate OCI root manifest
	manifest, manifestDescriptor, err := content.GenerateManifest(&configDescriptor, manifestAnnotations, layers...)
	if err != nil {
		logger.WithError(err).Fatal("[oci] failed to generate manfiest")
		return nil, err
	}

	// Store manifet to root
	err = store.StoreManifest(ref, manifestDescriptor, manifest)
	if err != nil {
		logger.WithError(err).Fatal("[oci] failed to store manfiest")
		return nil, err
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

	// Push to registry.
	// TODO: needs to verify root manifest digest
	rootManifest, err := oras.Copy(ctx, store, ref, registry, "", pushOpts...)
	if err != nil {
		logger.WithError(err).Fatal("[oci] failed push wasm module to registry")
		return nil, err
	}
	if rootManifest.Digest != manifestDescriptor.Digest {
		logger.WithError(err).Fatal("[oci] something error when verify your module, even it has been pushed successfully")
		return nil, err
	}

	if c.options.handleManifestDescriptor != nil {
		c.options.handleManifestDescriptor(rootManifest)
	}
	if c.options.handleLayerDescriptor != nil {
		c.options.handleLayerDescriptor(coreWasmContentLayerDesc)
	}

	image := &Image{
		Name:        refName.Name(),
		Labels:      map[string]string{},
		Manifest:    rootManifest,
		Webassembly: coreWasmContentLayerDesc,
		Content:     moduleContent,
	}

	return image, nil
}
