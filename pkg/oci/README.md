# Webassembly Open Container Interface

> Inspired by:
> * [https://github.com/engineerd/wasm-to-oci](https://github.com/engineerd/wasm-to-oci)
> * [https://github.com/oras-project/oras-go](https://github.com/oras-project/oras-go)

Push/Pull webassembly from registry. OCI manfiest sample json:

```json
{
  "schemaVersion": 2,
  "config": {
    "mediaType": "application/vnd.wasm.config.v1+json",
    "digest": "sha256:44136fa355b3678a1146ad16f7e8649e94fb4fc21fe77e8310c060f61caaff8a",
    "size": 2,
    "annotations": {
      "org.wasmerd.wasm.title": "example.wasm",
      "org.opencontainers.image.created": "2021-11-12T19:22:07+08:00"
    }
  },
  "layers": [
    {
      "mediaType": "application/vnd.wasm.content.layer.v1+wasm",
      "digest": "sha256:821d0b0539620fb5c9a3a0d2ef69c5bcb10c7f7c4b27b8b958882969c84a4695",
      "size": 1865223,
      "annotations": { "org.wasmerd.wasm.title": "example.wasm" }
    }
  ],
  "annotations": {
    "org.wasmerd.wasm.title": "example.wasm",
    "org.opencontainers.image.created": "2021-11-12T19:22:07+08:00"
  }
}
```


## Push example

```golang
client := oci.NewClient()
err := client.PushFromFile("ghcr.io/jeasonstudio/example-new.wasm:latest", "example.wasm")
```

## Pull example

```golang
client := oci.NewClient()
client.PullToFile("ghcr.io/jeasonstudio/example.wasm:latest", "target.wasm")
```
