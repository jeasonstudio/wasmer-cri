package main

import (
	"log"

	wasmercri "github.com/jeasonstudio/wasmer-cri"
)

var (
	config = &wasmercri.Config{
		Network: "unix",
		Address: "/tmp/wasmshim.sock",
	}
)

func main() {
	service, err := wasmercri.NewService(config)
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
