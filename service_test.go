package wasmercri

import (
	"log"
	"testing"
)

var (
	config = &Config{
		Network: "unix",
		Address: "/tmp/wasmshim.sock",
	}
)

func TestService(t *testing.T) {
	service, err := NewService(config)
	if err != nil {
		log.Fatalf("Init failed: %v", err)
		return
	}
	err = service.Listen()
	if err != nil {
		log.Fatalf("Init failed: %v", err)
		return
	}
}

// func TestWasm(t *testing.T) {
// 	if err := oci.Pull("wasm.xx", "local.wasm"); err != nil {
// 		log.Fatalf("cannot pull module: %v", err)
// 	}
// }
