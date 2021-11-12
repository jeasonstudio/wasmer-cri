package oci

// Thanks to https://github.com/engineerd/wasm-to-oci
const (
	// MediaTypeWasmContentLayer content layer media type
	MediaTypeWasmContentLayer = "application/vnd.wasm.content.layer.v1+wasm"
	// MediaTypeWasmConfig config media type
	MediaTypeWasmConfig = "application/vnd.wasm.config.v1+json"
	// AnnotationWasmTitle filename of webassembly, such as example.wasm
	AnnotationWasmTitle = "org.wasmerd.wasm.title"
)
