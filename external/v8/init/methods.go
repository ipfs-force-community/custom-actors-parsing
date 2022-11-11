package init

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs-force-community/custom-actors-parsing/external"
)

var Methods = map[uint64]external.MethodMeta{
	1: {"Constructor", *new(func(*ConstructorParams) abi.EmptyValue)},
	2: {"Exec", *new(func(*ExecParams) *ExecReturn)},
	3: {"Exec4", *new(func(*Exec4Params) *Exec4Return)},
	4: {"InstallCode", *new(func(*InstallParams) *InstallReturn)},
}
