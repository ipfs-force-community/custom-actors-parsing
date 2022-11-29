package paych

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v10/paych"
)

type UpdateChannelStateParams struct {
	Sv     paych.SignedVoucher
	Secret []byte
}

type ConstructorParams struct {
	From address.Address // Payer
	To   address.Address // Payee
}

// Modular Verification method
type ModVerifyParams struct {
	// Actor on which to invoke the method.
	Actor address.Address
	// Method to invoke.
	Method abi.MethodNum
	// Pre-serialized method parameters.
	Data []byte
}
