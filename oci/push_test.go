package oci

import (
	"testing"
	"time"
)

func TestPushWasm(t *testing.T) {
	err := Push("ghcr.io/jeasonstudio/example:v1", "example.wasm", nil)
	t.Log(err)
}

func TestTimeFormat(t *testing.T) {
	t.Log(time.Now().Format(time.RFC3339))
}
