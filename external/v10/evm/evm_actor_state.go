package evm

import "github.com/ipfs/go-cid"

type State struct {
	ByteCode      cid.Cid
	ContractState cid.Cid
	Nonce         uint64
}
