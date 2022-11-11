package init

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/ipfs/go-cid"
)

type State struct {
	AddressMap      cid.Cid
	NextID          abi.ActorID
	NetworkName     string
	InstalledActors cid.Cid
}
