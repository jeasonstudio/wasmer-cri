package oci

import (
	"testing"
)

func TestPullWasm(t *testing.T) {
	client := NewClient()
	client.PullToFile("ghcr.io/jeasonstudio/example.wasm:latest", "target.wasm")
}

// func TestImageDigest(t *testing.T) {
// 	dset := digestset.NewSet()
// 	desc, contents, _ := Pull("ghcr.io/jeasonstudio/example.wasm:latest", nil)
// 	t.Log(desc.Digest)

// 	image := digest.FromBytes(contents)
// 	err := image.Validate()
// 	t.Log(err)

// 	err = dset.Add(image)
// 	t.Log(err)

// 	newD, err := dset.Lookup(image.String())
// 	t.Log(newD)
// 	// image, _ := digest.Parse("sha256:E58FCF7418D4390DEC8E8FB69D88C06EC07039D651FEDD3AA72AF9972E7D046B")
// 	t.Log(err)
// }
