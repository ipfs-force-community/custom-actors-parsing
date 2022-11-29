package paych

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs-force-community/custom-actors-parsing/external"
)

var Methods = map[uint64]external.MethodMeta{
	1: {"Constructor", *new(func(*ConstructorParams) *abi.EmptyValue)},               // Constructor
	2: {"UpdateChannelState", *new(func(*UpdateChannelStateParams) *abi.EmptyValue)}, // UpdateChannelState
	3: {"Settle", *new(func(*abi.EmptyValue) *abi.EmptyValue)},                       // Settle
	4: {"Collect", *new(func(*abi.EmptyValue) *abi.EmptyValue)},                      // Collect
}
