package init

import (
	"github.com/filecoin-project/go-address"
	"github.com/ipfs/go-cid"
)

type ConstructorParams struct {
	NetworkName string
}

type ExecParams struct {
	CodeCid           cid.Cid
	ConstructorParams []byte
}

type ExecReturn struct {
	IdAddress     address.Address
	RobustAddress address.Address
}

type Exec4Params struct {
	CodeCid           cid.Cid
	ConstructorParams []byte
	SubAddress        []byte
}

type Exec4Return = ExecReturn

type InstallParams struct {
	Code []byte
}

type InstallReturn struct {
	CodeCid   cid.Cid
	Installed bool
}
