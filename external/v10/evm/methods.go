package evm

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin"
	"github.com/holiman/uint256"
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

var Methods = map[abi.MethodNum]builtin.MethodMeta{
	1: {"Constructor", *new(func(*ConstructorParams) *abi.EmptyValue)},
	// todo: params to *    注册每个独立evm actor的方法
	2: {"InvokeContract", *new(func(InputData []byte, WithCode cid.Cid) []byte)},
	3: {"GetBytecode", *new(func(*abi.EmptyValue) cid.Cid)},
	4: {"GetStorageAt", *new(func(*GetStorageAtParams) uint256.Int)},
	// todo
	5: {"InvokeContractDelegate", *new(func(InputData []byte, WithCode cid.Cid) []byte)},
	6: {"GetBytecodeHash", *new(func(*abi.EmptyValue) multihash.Multihash)},
}
