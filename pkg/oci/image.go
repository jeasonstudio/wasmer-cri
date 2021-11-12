package oci

import (
	"time"

	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// Image is the OCI Image
type Image struct {
	// Name of the webassembly image.
	//
	// To be pulled, it must be a reference compatible with resolvers.
	Name string

	// Labels provide runtime decoration for the image record.
	//
	// There is no default behavior for how these labels are propagated. They
	// only decorate the static metadata object.
	// This field is optional.
	Labels map[string]string

	// Manifest describes the root content for this image. Typically, this is
	// a manifest, index or manifest list.
	Manifest v1.Descriptor

	// Webassembly describes the webassembly layer content for this image.
	Webassembly v1.Descriptor

	// Content describes webassembly file content.
	Content []byte

	CreatedAt, UpdatedAt time.Time
}

// Size get webassembly file size
func (i *Image) Size() int64 {
	return i.Webassembly.Size
}
