package multisig

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	"github.com/ipfs/go-cid"
)

type State struct {
	Signers               []address.Address // Signers must be canonical ID-addresses.
	NumApprovalsThreshold uint64
	NextTxnID             multisig.TxnID

	// Linear unlock
	InitialBalance abi.TokenAmount
	StartEpoch     abi.ChainEpoch
	UnlockDuration abi.ChainEpoch

	PendingTxns cid.Cid // HAMT[TxnID]Transaction
}
