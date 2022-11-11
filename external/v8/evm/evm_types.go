package evm

import (
	"github.com/holiman/uint256"
)

type ConstructorParams struct {
	Creator  [20]byte
	InitCode []byte
}

type GetStorageAtParams struct {
	StorageKey uint256.Int
}
