package verifreg

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v8/verifreg"
)

type AddVerifierParams struct {
	Address   address.Address
	Allowance verifreg.DataCap
}

type AddVerifiedClientParams struct {
	Address   address.Address
	Allowance verifreg.DataCap
}

type UseBytesParams struct {
	Address  address.Address  // Address of verified client.
	DealSize abi.StoragePower // Number of bytes to use.
}

type RestoreBytesParams struct {
	Address  address.Address
	DealSize abi.StoragePower
}

type RemoveDataCapParams struct {
	VerifiedClientToRemove address.Address
	DataCapAmountToRemove  verifreg.DataCap
	VerifierRequest1       verifreg.RemoveDataCapRequest
	VerifierRequest2       verifreg.RemoveDataCapRequest
}

type RemoveDataCapReturn struct {
	VerifiedClient address.Address
	DataCapRemoved verifreg.DataCap
}
