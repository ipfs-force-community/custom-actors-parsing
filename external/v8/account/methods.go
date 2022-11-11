package account

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs-force-community/custom-actors-parsing/external"
)

var Methods = map[uint64]external.MethodMeta{
	1: {"Constructor", *new(func(*address.Address) *abi.EmptyValue)},
	2: {"PubkeyAddress", *new(func(*abi.EmptyValue) *address.Address)},
	3: {"AuthenticateMessage", *new(func(*AuthenticateMessageParams) *abi.EmptyValue)},
}
