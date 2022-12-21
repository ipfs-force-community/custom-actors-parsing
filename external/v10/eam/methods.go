package eam

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
)

var Methods = map[abi.MethodNum]builtin.MethodMeta{
	1: {"Constructor", *new(func(value *abi.EmptyValue) *abi.EmptyValue)},
	2: {"Create", *new(func(*CreateParams) *CreateReturn)},
	3: {"Create2", *new(func(*Create2Params) *Create2Return)},
}
