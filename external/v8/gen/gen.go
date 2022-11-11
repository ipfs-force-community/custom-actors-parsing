package main

import (
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/account"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/eam"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/evm"
	init_ "github.com/ipfs-force-community/custom-actors-parsing/external/v8/init"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/market"

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

	if err := gen.WriteTupleEncodersToFile("./external/v8/init/cbor_gen.go", "init",
		init_.ConstructorParams{},
		init_.ExecParams{},
		init_.ExecReturn{},
		init_.Exec4Params{},
		init_.InstallParams{},
		init_.InstallReturn{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/account/cbor_gen.go", "account",
		account.AuthenticateMessageParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/market/cbor_gen.go", "market",
		market.WithdrawBalanceParams{},
		market.WithdrawBalanceReturn{},
		market.PublishStorageDealsParams{},
		market.PublishStorageDealsReturn{},
		market.VerifyDealsForActivationParams{},
		market.VerifyDealsForActivationReturn{},
		market.ActivateDealsParams{},
		market.OnMinerSectorsTerminateParams{},
		market.ComputeDataCommitmentParams{},
		market.ComputeDataCommitmentReturn{},
		market.DealHealthyParams{},
	); err != nil {
		panic(err)
	}
}
