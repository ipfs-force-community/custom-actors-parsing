package evm

import (
	"custom-actors-parsing/external"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/holiman/uint256"
	"github.com/ipfs/go-cid"
)

var Methods = map[uint64]external.MethodMeta{
	1: {"Constructor", *new(func(*ConstructorParams) *abi.EmptyValue)},
	// todo: params to *    注册每个独立evm actor的方法
	2: {"InvokeContract", *new(func(Method uint64, InputData []byte, ReadOnly bool, WithCode cid.Cid) []byte)},
	3: {"GetBytecode", *new(func(*abi.EmptyValue) cid.Cid)},
	4: {"GetStorageAt", *new(func(*GetStorageAtParams) uint256.Int)},
	// todo
	5: {"InvokeContractReadOnly", *new(func(Method uint64, InputData []byte, ReadOnly bool, WithCode cid.Cid) []byte)},
	6: {"InvokeContractDelegate", *new(func(Method uint64, InputData []byte, ReadOnly bool, WithCode cid.Cid) []byte)},
}
