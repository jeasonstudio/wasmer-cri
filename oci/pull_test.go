package oci

import (
	"testing"
)

func TestPullWasm(t *testing.T) {
	desc, _ := Pull("ghcr.io/jeasonstudio/example:v1", "target.wasm", nil)
	t.Log(desc.Digest)
	t.Log(desc.Size)
	t.Log(desc.Annotations)
}
