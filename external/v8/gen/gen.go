package main

import (
	"custom-actors-parsing/external/v8/eam"
	"custom-actors-parsing/external/v8/evm"

	gen "github.com/whyrusleeping/cbor-gen"
)

func main() {
	if err := gen.WriteTupleEncodersToFile("./external/v8/evm/cbor_gen.go", "evm",
		evm.State{},
		evm.ConstructorParams{},
		evm.GetStorageAtParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/eam/cbor_gen.go", "eam",
		eam.CreateParams{},
		eam.CreateReturn{},
		eam.Create2Params{},
		eam.Create2Return{},
	); err != nil {
		panic(err)
	}
}
