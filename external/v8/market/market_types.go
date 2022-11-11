package market

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-bitfield"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v8/market"
	"github.com/ipfs/go-cid"
)

type WithdrawBalanceParams struct {
	ProviderOrClient address.Address
	Amount           abi.TokenAmount
}

type WithdrawBalanceReturn struct {
	AmountWithdrawn abi.TokenAmount
}

type PublishStorageDealsParams struct {
	Deals []market.ClientDealProposal
}

type PublishStorageDealsReturn struct {
	IDs        []abi.DealID
	ValidDeals bitfield.BitField
}

type VerifyDealsForActivationParams struct {
	Sectors []market.SectorDeals
}

type VerifyDealsForActivationReturn struct {
	Sectors []market.SectorWeights
}

type ActivateDealsParams struct {
	DealIDs      []abi.DealID
	SectorExpiry abi.ChainEpoch
}

type OnMinerSectorsTerminateParams struct {
	Epoch   abi.ChainEpoch
	DealIDs []abi.DealID
}

type ComputeDataCommitmentParams struct {
	Inputs []market.SectorDataSpec
}

type ComputeDataCommitmentReturn struct {
	Commds []cid.Cid
}

type DealHealthyParams struct {
	Piece cid.Cid
	ID    abi.DealID
}
