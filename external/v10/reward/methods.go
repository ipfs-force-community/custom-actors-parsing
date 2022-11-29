package reward

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs-force-community/custom-actors-parsing/external"
)

var Methods = map[uint64]external.MethodMeta{
	1: {"Constructor", *new(func(*abi.StoragePower) *abi.EmptyValue)},            // Constructor
	2: {"AwardBlockReward", *new(func(*AwardBlockRewardParams) *abi.EmptyValue)}, // AwardBlockReward
	3: {"ThisEpochReward", *new(func(*abi.EmptyValue) *ThisEpochRewardReturn)},   // ThisEpochReward
	4: {"UpdateNetworkKPI", *new(func(*abi.StoragePower) *abi.EmptyValue)},       // UpdateNetworkKPI
}
