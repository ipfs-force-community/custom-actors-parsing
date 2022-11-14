package main

import (
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/account"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/cron"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/eam"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/evm"
	init_ "github.com/ipfs-force-community/custom-actors-parsing/external/v8/init"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/market"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/miner"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/multisig"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/paych"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/power"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/reward"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/system"
	"github.com/ipfs-force-community/custom-actors-parsing/external/v8/verifreg"

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
		init_.State{},
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
		account.State{},
		account.AuthenticateMessageParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/market/cbor_gen.go", "market",
		market.State{},
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

	if err := gen.WriteTupleEncodersToFile("./external/v8/cron/cbor_gen.go", "cron",
		cron.State{},
		cron.ConstructorParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/miner/cbor_gen.go", "miner",
		// actor state
		miner.State{},
		// method params and returns
		miner.GetControlAddressesReturn{},
		miner.ChangeWorkerAddressParams{},
		miner.ChangePeerIDParams{},
		miner.SubmitWindowedPoStParams{},
		miner.PreCommitSectorParams{},
		miner.ProveCommitSectorParams{},
		miner.ExtendSectorExpirationParams{},
		miner.TerminateSectorsParams{},
		miner.TerminateSectorsReturn{},
		miner.DeclareFaultsParams{},
		miner.DeclareFaultsRecoveredParams{},
		miner.DeferredCronEventParams{},
		miner.CheckSectorProvenParams{},
		miner.ApplyRewardParams{},
		miner.ReportConsensusFaultParams{},
		miner.WithdrawBalanceParams{},
		miner.ConfirmSectorProofsParams{},
		miner.ChangeMultiaddrsParams{},
		miner.CompactPartitionsParams{},
		miner.CompactSectorNumbersParams{},
		miner.DisputeWindowedPoStParams{},
		miner.PreCommitSectorBatchParams{},
		miner.ProveCommitAggregateParams{},
		miner.ProveReplicaUpdatesParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/multisig/cbor_gen.go", "multisig",
		multisig.State{},
		multisig.ConstructorParams{},
		multisig.ProposeParams{},
		multisig.ProposeReturn{},
		multisig.AddSignerParams{},
		multisig.RemoveSignerParams{},
		multisig.TxnIDParams{},
		multisig.ApproveReturn{},
		multisig.ChangeNumApprovalsThresholdParams{},
		multisig.SwapSignerParams{},
		multisig.LockBalanceParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/paych/cbor_gen.go", "paych",
		// actor state
		paych.State{},
		//method params and returns
		paych.ConstructorParams{},
		paych.UpdateChannelStateParams{},
		paych.ModVerifyParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/power/cbor_gen.go", "power",
		// actors state
		power.State{},
		// method params and returns
		power.UpdateClaimedPowerParams{},
		power.MinerConstructorParams{},
		power.CreateMinerParams{},
		power.CreateMinerReturn{},
		power.CurrentTotalPowerReturn{},
		power.EnrollCronEventParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/reward/cbor_gen.go", "reward",
		// actor state
		reward.State{},
		// method params and returns
		reward.ThisEpochRewardReturn{},
		reward.AwardBlockRewardParams{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/system/cbor_gen.go", "system",
		// actor state
		system.State{},
	); err != nil {
		panic(err)
	}

	if err := gen.WriteTupleEncodersToFile("./external/v8/verifreg/cbor_gen.go", "verifreg",
		// actor state
		verifreg.State{},
		// method params and returns
		verifreg.AddVerifierParams{},
		verifreg.AddVerifiedClientParams{},
		verifreg.UseBytesParams{},
		verifreg.RestoreBytesParams{},
		verifreg.RemoveDataCapParams{},
		verifreg.RemoveDataCapReturn{},
	); err != nil {
		panic(err)
	}

}
