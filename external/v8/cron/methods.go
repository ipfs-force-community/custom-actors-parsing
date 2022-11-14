package cron

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs-force-community/custom-actors-parsing/external"
)

var Methods = map[uint64]external.MethodMeta{
	1: {"Constructor", *new(func(*ConstructorParams) *abi.EmptyValue)},
	2: {"EpochTick", *new(func(*abi.EmptyValue) *address.Address)},
}
