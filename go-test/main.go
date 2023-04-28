package main

import (
	"fmt"
	"io/ioutil"

	wasmer "github.com/wasmerio/wasmer-go/wasmer"
)

func main() {
	wasmBytes, _ := ioutil.ReadFile("wapm_packages/enricopdg/tutorial-1@0.1.0/tutorial_1.wasm")

	engine := wasmer.NewEngine()
	store := wasmer.NewStore(engine)

	module, _ := wasmer.NewModule(store, wasmBytes)

	importObject := wasmer.NewImportObject()
	instance, _ := wasmer.NewInstance(module, importObject)

	add, _ := instance.Exports.GetFunction("add")

	var a, b float32 = 1.0, 2.0

	result, err := add(a, b)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
