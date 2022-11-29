package paych

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

type State struct {
	// Channel owner, who has funded the actor
	From address.Address
	// Recipient of payouts from channel
	To address.Address

	// Amount successfully redeemed through the payment channel, paid out on `Collect()`
	ToSend abi.TokenAmount

	// Height at which the channel can be `Collected`
	SettlingAt abi.ChainEpoch
	// Height before which the channel `ToSend` cannot be collected
	MinSettleHeight abi.ChainEpoch

	// Collections of lane states for the channel, maintained in ID order.
	LaneStates cid.Cid // AMT<LaneState>
}
