package main

import (
	"fmt"
	"io/ioutil"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	wasmBytes, _ := ioutil.ReadFile("simple.wasm")
	// wasmBytes, _ := ioutil.ReadFile("./target/wasm32-wasi/debug/example.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, _ := wasmer.NewModule(store, wasmBytes)

	wasiEnv, _ := wasmer.NewWasiStateBuilder("example").Argument("--foo").
		Environment("ABC", "DEF").
		Environment("X", "ZY").
		MapDirectory("the_host_current_directory", ".").
		CaptureStdout().
		Finalize()

	// Instantiates the module
	importObject, _ := wasiEnv.GenerateImportObject(store, module)

	instance, _ := wasmer.NewInstance(module, importObject)

	// Gets the `sum` exported function from the WebAssembly instance.
	// sum, _ := instance.Exports.GetFunction("sum")

	myMain, _ := instance.Exports.GetWasiStartFunction()

	res, err := myMain()
	fmt.Println(err)

	// Calls that exported function with Go standard values. The WebAssembly
	// types are inferred and values are casted automatically.
	// result, _ := sum(5, 37)

	fmt.Println(res) // 42!
	// fmt.Println(result) // 42!
}
