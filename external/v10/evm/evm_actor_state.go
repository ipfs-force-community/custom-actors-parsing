package evm

import (
	"github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

type State struct {
	ByteCode      cid.Cid
	ByteCodeHash  multihash.Multihash
	ContractState cid.Cid
	Nonce         uint64
}
