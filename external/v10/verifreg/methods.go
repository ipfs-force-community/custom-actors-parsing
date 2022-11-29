package verifreg

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs-force-community/custom-actors-parsing/external"

	"github.com/filecoin-project/go-state-types/abi"
)

var Methods = map[uint64]external.MethodMeta{
	1: {"Constructor", *new(func(*address.Address) *abi.EmptyValue)},                          // Constructor
	2: {"AddVerifier", *new(func(*AddVerifierParams) *abi.EmptyValue)},                        // AddVerifier
	3: {"RemoveVerifier", *new(func(*address.Address) *abi.EmptyValue)},                       // RemoveVerifier
	4: {"AddVerifiedClient", *new(func(*AddVerifiedClientParams) *abi.EmptyValue)},            // AddVerifiedClient
	5: {"UseBytes", *new(func(*UseBytesParams) *abi.EmptyValue)},                              // UseBytes
	6: {"RestoreBytes", *new(func(*RestoreBytesParams) *abi.EmptyValue)},                      // RestoreBytes
	7: {"RemoveVerifiedClientDataCap", *new(func(*RemoveDataCapParams) *RemoveDataCapReturn)}, // RemoveVerifiedClientDataCap
}
