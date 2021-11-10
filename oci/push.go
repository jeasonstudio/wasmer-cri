package oci

import (
	"context"
	"io/ioutil"
	"path/filepath"
	"time"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
	log "github.com/sirupsen/logrus"
	"oras.land/oras-go/pkg/content"
	"oras.land/oras-go/pkg/oras"
)

// Push pushes a WASM module to an OCI registry
func Push(ref, filename string, moduleContent []byte, registryConfig *RegistryConfig) error {
	log.WithFields(log.Fields{
		"ref":            ref,
		"registryConfig": registryConfig,
		"filename":       filename,
	}).Debug("PushFromFile")

	ctx := context.Background()
	store := content.NewMemory()

	descriptor, err := store.Add(filename, ConfigMediaType, moduleContent)
	if err != nil {
		log.WithError(err).Fatal("Failed to read Webassembly file")
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

// PushFromFile push to registry from file in file system
func PushFromFile(ref, modulePath string, registryConfig *RegistryConfig) error {
	filename := filepath.Base(modulePath)

	log.WithFields(log.Fields{
		"ref":            ref,
		"path":           modulePath,
		"registryConfig": registryConfig,
		"filename":       filename,
	}).Debug("PushFromFile")

	fileContent, err := ioutil.ReadFile(modulePath)
	if err != nil {
		log.WithError(err).Fatal("Failed to read Webassembly file from file-system")
		return err
	}

	err = Push(ref, filename, fileContent, registryConfig)
	if err != nil {
		log.WithError(err).Fatal("Failed to push Webassembly")
		return err
	}
	return nil
}
