package main

import (
	"fmt"
	"io/ioutil"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	// wasmBytes, _ := ioutil.ReadFile("simple.wasm")
	wasmBytes, _ := ioutil.ReadFile("./target/wasm32-wasi/debug/example.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	// Compiles the module
	module, _ := wasmer.NewModule(store, wasmBytes)

	wasiEnv, _ := wasmer.NewWasiStateBuilder("example").Argument("/hello.txt").Argument("asdf").MapDirectory("/", ".").InheritStdout().Finalize()

	// Instantiates the module
	importObject, _ := wasiEnv.GenerateImportObject(store, module)

	hostFunction := wasmer.NewFunction(
		store,
		wasmer.NewFunctionType(wasmer.NewValueTypes(), wasmer.NewValueTypes(wasmer.I32)),
		func(args []wasmer.Value) ([]wasmer.Value, error) {
			return []wasmer.Value{wasmer.NewI32(42)}, nil
		},
	)
	importObject.Register(
		"hello",
		map[string]wasmer.IntoExtern{
			"host_function": hostFunction,
		},
	)

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
