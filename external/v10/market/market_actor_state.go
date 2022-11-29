package market

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

type State struct {
	Proposals                     cid.Cid
	States                        cid.Cid
	PendingProposals              cid.Cid
	EscrowTable                   cid.Cid
	LockedTable                   cid.Cid
	NextID                        abi.DealID
	DealOpsByEpoch                cid.Cid
	LastCron                      abi.ChainEpoch
	TotalClientLockedCollateral   abi.TokenAmount
	TotalProviderLockedCollateral abi.TokenAmount
	TotalClientStorageFee         abi.TokenAmount
}
