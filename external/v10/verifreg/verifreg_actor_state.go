package verifreg

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
)

type State struct {
	RootKey                  address.Address
	Verifiers                cid.Cid // HAMT[addr.Address]DataCap
	VerifiedClients          cid.Cid // HAMT[addr.Address]DataCap
	RemoveDataCapProposalIDs cid.Cid // HAMT[AddrPairKey]RmDcProposalID
}
