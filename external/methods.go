package external

import "github.com/filecoin-project/go-state-types/abi"

const (
	MethodSend                     = abi.MethodNum(0)
	MethodConstructor              = abi.MethodNum(1)
	UniversalReceiverHookMethodNum = abi.MethodNum(3726118371)
)

var MethodsEam = struct {
	Constructor abi.MethodNum
	Create      abi.MethodNum
	Create2     abi.MethodNum
}{MethodConstructor, 2, 3}

var MethodsEvm = struct {
	Constructor            abi.MethodNum
	InvokeContract         abi.MethodNum
	GetBytecode            abi.MethodNum
	GetStorageAt           abi.MethodNum
	InvokeContractReadOnly abi.MethodNum
	InvokeContractDelegate abi.MethodNum
}{MethodConstructor, 2, 3, 4, 5, 6}
