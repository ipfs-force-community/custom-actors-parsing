package multisig

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v10/multisig"
	"github.com/filecoin-project/go-state-types/exitcode"
)

type ConstructorParams struct {
	Signers               []address.Address
	NumApprovalsThreshold uint64
	UnlockDuration        abi.ChainEpoch
	StartEpoch            abi.ChainEpoch
}

type ProposeParams struct {
	To     address.Address
	Value  abi.TokenAmount
	Method abi.MethodNum
	Params []byte
}

type ProposeReturn struct {
	// TxnID is the ID of the proposed transaction
	TxnID multisig.TxnID
	// Applied indicates if the transaction was applied as opposed to proposed but not applied due to lack of approvals
	Applied bool
	// Code is the exitcode of the transaction, if Applied is false this field should be ignored.
	Code exitcode.ExitCode
	// Ret is the return vale of the transaction, if Applied is false this field should be ignored.
	Ret []byte
}

type TxnIDParams struct {
	ID multisig.TxnID
	// Optional hash of proposal to ensure an operation can only apply to a
	// specific proposal.
	ProposalHash []byte
}

type ApproveReturn struct {
	// Applied indicates if the transaction was applied as opposed to proposed but not applied due to lack of approvals
	Applied bool
	// Code is the exitcode of the transaction, if Applied is false this field should be ignored.
	Code exitcode.ExitCode
	// Ret is the return vale of the transaction, if Applied is false this field should be ignored.
	Ret []byte
}

type AddSignerParams struct {
	Signer   address.Address
	Increase bool
}

type RemoveSignerParams struct {
	Signer   address.Address
	Decrease bool
}

type SwapSignerParams struct {
	From address.Address
	To   address.Address
}

type ChangeNumApprovalsThresholdParams struct {
	NewThreshold uint64
}

type LockBalanceParams struct {
	StartEpoch     abi.ChainEpoch
	UnlockDuration abi.ChainEpoch
	Amount         abi.TokenAmount
}
