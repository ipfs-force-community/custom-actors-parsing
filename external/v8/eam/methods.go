package eam

import (
	"custom-actors-parsing/external"

	"github.com/filecoin-project/go-state-types/abi"
)

var Methods = map[uint64]external.MethodMeta{
	1: {"Constructor", *new(func(value *abi.EmptyValue) *abi.EmptyValue)},
	2: {"Create", *new(func(*CreateParams) *CreateReturn)},
	3: {"Create2", *new(func(*Create2Params) *Create2Return)},
}
