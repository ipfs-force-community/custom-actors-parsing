package miner

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v8/miner"
	"github.com/filecoin-project/go-state-types/builtin/v8/power"
	"github.com/filecoin-project/go-state-types/builtin/v8/util/smoothing"
	"github.com/filecoin-project/go-state-types/proof"
	"github.com/ipfs/go-cid"
)

type DeclareFaultsRecoveredParams struct {
	Recoveries []miner.RecoveryDeclaration
}

type DeclareFaultsParams struct {
	Faults []miner.FaultDeclaration
}

type ProveReplicaUpdatesParams struct {
	Updates []miner.ReplicaUpdate
}

type SubmitWindowedPoStParams struct {
	// The deadline index which the submission targets.
	Deadline uint64
	// The partitions being proven.
	Partitions []miner.PoStPartition
	// Array of proofs, one per distinct registered proof type present in the sectors being proven.
	// In the usual case of a single proof type, this array will always have a single element (independent of number of partitions).
	Proofs []proof.PoStProof
	// The epoch at which these proofs is being committed to a particular chain.
	ChainCommitEpoch abi.ChainEpoch
	// The ticket randomness on the chain at the ChainCommitEpoch on the chain this post is committed to
	ChainCommitRand abi.Randomness
}

type DisputeWindowedPoStParams struct {
	Deadline  uint64
	PoStIndex uint64 // only one is allowed at a time to avoid loading too many sector infos.
}

type ProveCommitAggregateParams struct {
	SectorNumbers  bitfield.BitField
	AggregateProof []byte
}

type ProveCommitSectorParams struct {
	SectorNumber abi.SectorNumber
	Proof        []byte
}

type MinerConstructorParams = power.MinerConstructorParams

type TerminateSectorsParams struct {
	Terminations []miner.TerminationDeclaration
}

type TerminateSectorsReturn struct {
	// Set to true if all early termination work has been completed. When
	// false, the miner may choose to repeatedly invoke TerminateSectors
	// with no new sectors to process the remainder of the pending
	// terminations. While pending terminations are outstanding, the miner
	// will not be able to withdraw funds.
	Done bool
}

type ChangePeerIDParams struct {
	NewID abi.PeerID
}

type ChangeMultiaddrsParams struct {
	NewMultiaddrs []abi.Multiaddrs
}

type ChangeWorkerAddressParams struct {
	NewWorker       address.Address
	NewControlAddrs []address.Address
}

type ExtendSectorExpirationParams struct {
	Extensions []miner.ExpirationExtension
}

type ReportConsensusFaultParams struct {
	BlockHeader1     []byte
	BlockHeader2     []byte
	BlockHeaderExtra []byte
}

type GetControlAddressesReturn struct {
	Owner        address.Address
	Worker       address.Address
	ControlAddrs []address.Address
}

type CheckSectorProvenParams struct {
	SectorNumber abi.SectorNumber
}

type WithdrawBalanceParams struct {
	AmountRequested abi.TokenAmount
}

type CompactPartitionsParams struct {
	Deadline   uint64
	Partitions bitfield.BitField
}

type CompactSectorNumbersParams struct {
	MaskSectorNumbers bitfield.BitField
}

type PreCommitSectorBatchParams struct {
	Sectors []miner.SectorPreCommitInfo
}

type PreCommitSectorParams struct {
	SealProof       abi.RegisteredSealProof
	SectorNumber    abi.SectorNumber
	SealedCID       cid.Cid `checked:"true"` // CommR
	SealRandEpoch   abi.ChainEpoch
	DealIDs         []abi.DealID
	Expiration      abi.ChainEpoch
	ReplaceCapacity bool // DEPRECATED: Whether to replace a "committed capacity" no-deal sector (requires non-empty DealIDs)
	// DEPRECATED: The committed capacity sector to replace, and it's deadline/partition location
	ReplaceSectorDeadline  uint64
	ReplaceSectorPartition uint64
	ReplaceSectorNumber    abi.SectorNumber
}

type DeferredCronEventParams struct {
	EventPayload            []byte
	RewardSmoothed          smoothing.FilterEstimate
	QualityAdjPowerSmoothed smoothing.FilterEstimate
}

type ApplyRewardParams struct {
	Reward  abi.TokenAmount
	Penalty abi.TokenAmount
}

type ConfirmSectorProofsParams struct {
	Sectors                 []abi.SectorNumber
	RewardSmoothed          smoothing.FilterEstimate
	RewardBaselinePower     abi.StoragePower
	QualityAdjPowerSmoothed smoothing.FilterEstimate
}
