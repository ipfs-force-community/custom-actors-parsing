package parse

import (
	"encoding/json"
	"fmt"
	"github.com/ipfs-force-community/custom-actors-parsing/registry"
	"unsafe"
)

func ParseInvokeContractParams(params []byte) ([]byte, error) {
	// parse contract method
	hexParams, err := registry.HexEncodeByteArray(params)
	if err != nil {
		return nil, fmt.Errorf("hex encode params failed: %w", err)
	}
	hexParamsString := *(*string)(unsafe.Pointer(&hexParams))

	// the first 4 bytes is method ID
	if len(hexParamsString) < 8 {
		return nil, fmt.Errorf("invalid length of params %v for InvokeContract", len(hexParams))
	}
	if (len(hexParamsString)-8)%64 != 0 {
		return nil, fmt.Errorf("invalid length of params %v for InvokeContract", len(hexParams))
	}

	methodID := hexParamsString[:8]
	datas := hexParamsString[8:]

	inputData, err := registry.GetInputDataByMethodID(fmt.Sprintf("%s%s", "0x", methodID), datas)
	if err != nil {
		// todo: err just log instead of returning
		return nil, err
	}

	input, err := json.Marshal(inputData)
	if err != nil {
		return nil, err
	}

	return input, nil

}
