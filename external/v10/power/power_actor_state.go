package power

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v10/util/smoothing"
	"github.com/ipfs/go-cid"
)

type State struct {
	TotalRawBytePower         abi.StoragePower
	TotalBytesCommitted       abi.StoragePower
	TotalQualityAdjPower      abi.StoragePower
	TotalQABytesCommitted     abi.StoragePower
	TotalPledgeCollateral     abi.TokenAmount
	ThisEpochRawBytePower     abi.StoragePower
	ThisEpochQualityAdjPower  abi.StoragePower
	ThisEpochPledgeCollateral abi.TokenAmount
	ThisEpochQAPowerSmoothed  smoothing.FilterEstimate
	MinerCount                int64
	MinerAboveMinPowerCount   int64
	CronEventQueue            cid.Cid // Multimap, (HAMT[ChainEpoch]AMT[CronEvent])
	FirstCronEpoch            abi.ChainEpoch
	Claims                    cid.Cid  // Map, HAMT[address]Claim
	ProofValidationBatch      *cid.Cid // Multimap, (HAMT[Address]AMT[SealVerifyInfo])
}

type Claim struct {
	// Miner's proof type used to determine minimum miner size
	WindowPoStProofType abi.RegisteredPoStProof

	// Sum of raw byte power for a miner's sectors.
	RawBytePower abi.StoragePower

	// Sum of quality adjusted power for a miner's sectors.
	QualityAdjPower abi.StoragePower
}
