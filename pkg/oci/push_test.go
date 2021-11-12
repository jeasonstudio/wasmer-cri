package oci

import (
	"testing"
	"time"
)

func TestPushWasm(t *testing.T) {
	client := NewClient()
	err := client.PushFromFile("ghcr.io/jeasonstudio/example-new.wasm:latest", "example.wasm")
	t.Log(err)
}

func TestTimeFormat(t *testing.T) {
	t.Log(time.Now().Format(time.RFC3339))
}

// func TestParseImageRef(t *testing.T) {
// 	ref, _ := docker.ParseAnyReference("ghcr.io/jeasonstudio/foo:latest")
// 	t.Log(ref.String())
// }
