package power

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v10/util/smoothing"
)

type MinerConstructorParams struct {
	OwnerAddr           address.Address
	WorkerAddr          address.Address
	ControlAddrs        []address.Address
	WindowPoStProofType abi.RegisteredPoStProof
	PeerId              abi.PeerID
	Multiaddrs          []abi.Multiaddrs
}

type CreateMinerParams struct {
	Owner               address.Address
	Worker              address.Address
	WindowPoStProofType abi.RegisteredPoStProof
	Peer                abi.PeerID
	Multiaddrs          []abi.Multiaddrs
}

type CreateMinerReturn struct {
	IDAddress     address.Address // The canonical ID-based address for the actor.
	RobustAddress address.Address // A more expensive but re-org-safe address for the newly created actor.
}

type UpdateClaimedPowerParams struct {
	RawByteDelta         abi.StoragePower
	QualityAdjustedDelta abi.StoragePower
}

type EnrollCronEventParams struct {
	EventEpoch abi.ChainEpoch
	Payload    []byte
}

type CurrentTotalPowerReturn struct {
	RawBytePower            abi.StoragePower
	QualityAdjPower         abi.StoragePower
	PledgeCollateral        abi.TokenAmount
	QualityAdjPowerSmoothed smoothing.FilterEstimate
}
