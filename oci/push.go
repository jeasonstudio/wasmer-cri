package oci

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
)

// Push pushes a WASM module to an OCI registry
func Push(ref, module string, registryConfig *RegistryConfig) error {
	ctx := context.Background()
	store := content.NewMemory()

	fileContent, err := ioutil.ReadFile(module)
	if err != nil {
		log.Fatalf("Fail to read Webassembly file, %s", err)
		return err
	}

	descriptor, err := store.Add(module, ConfigMediaType, fileContent)
	if err != nil {
		log.Fatalf("Fail to read Webassembly file, %s", err)
		return err
	}

	manifestAnnotations := map[string]string{
		v1.AnnotationAuthors:       "",
		v1.AnnotationCreated:       time.Now().Format(time.RFC3339),
		v1.AnnotationDescription:   "",
		v1.AnnotationDocumentation: "",
		v1.AnnotationLicenses:      "",
		v1.AnnotationVersion:       "1.0.0",
	}
	configAnnotations := map[string]string{}

	manifest, manifestDescriptor, config, configDescriptor, err := content.GenerateManifestAndConfig(manifestAnnotations, configAnnotations, descriptor)
	store.Set(configDescriptor, config)
	err = store.StoreManifest(ref, manifestDescriptor, manifest)
	if err != nil {
		log.Fatalf("Fail to push module %s", err)
		return err
	}

	registry, err := NewRegistry(registryConfig)
	if err != nil {
		log.Fatalf("Fail to push module, %s", err)
		return err
	}

	_, err = oras.Copy(ctx, store, ref, registry, "")
	if err != nil {
		log.Fatalf("Fail to push module, %s", err)
		return err
	}

	return nil
}
