package registry

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common/hexutil"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/crypto/sha3"
)

type ConstractMethodsRegistry struct {
	sync.RWMutex
	c map[string]*RegisteredInputData
}

// MethodsRegistry is used to registry smart contract code, query from db cronly & update when users registry operate
var MethodsRegistry = struct {
	methods *ConstractMethodsRegistry
}{
	methods: &ConstractMethodsRegistry{
		c: make(map[string]*RegisteredInputData),
	},
}

//// RegistryMethodID: method(params...)
//balanceOf(address):70a08231
//totalSupply(): 18160ddd
//transfer(address, uint256): 9d61d234
//load(): 86d5c4be
//withdraw(address,uint256,address): 69328dec

// get from db
const (
	balanceOf   = "balanceOf(address account)"
	totalSupply = "totalSupply()"
	withdraw    = "withdraw(address token, uint256 amount, address destination)"
)

// known
// todo: 返回类型  json abi
//func init() {
//	// get from db
//	functionList := []string{balanceOf, totalSupply, withdraw}
//
//	if err := RegisterConstractMethods(functionList); err != nil {
//		panic(err)
//	}
//}

func RegisterConstractMethods(functions []string) error {
	var (
		methodID string
		err      error
	)
	MethodsRegistry.methods.Lock()

	// lock内计算，会导致lock锁占有过长时间？？
	for _, function := range functions {
		inputData := &RegisteredInputData{}

		methodID, err = GetMethodID(function)
		if err != nil {
			return err
		}

		params, err := GetConstractParams(function)
		if err != nil {
			return err
		}

		inputData.MethodID = methodID
		inputData.Function = function
		inputData.Params = params
		MethodsRegistry.methods.c[methodID] = inputData
	}

	MethodsRegistry.methods.Unlock()

	// todo: 落库

	return nil
}

func GetMethodID(function string) (string, error) {
	constractParams, err := GetConstractParams(function)
	if err != nil {
		return "", err
	}

	var paramsList []byte
	for i, params := range constractParams {
		paramsList = append(paramsList, params.Type...)
		if i < len(constractParams)-1 {
			paramsList = append(paramsList, ',')
		}
	}

	var all []byte
	all = append(all, getFunctionName(function)...)
	all = append(all, '(')
	all = append(all, paramsList...)
	all = append(all, ')')

	var buffer bytes.Buffer
	_, err = buffer.Write(all)
	if err != nil {
		return "", err
	}

	hasher := sha3.NewLegacyKeccak256()
	hasher.Write(buffer.Bytes())
	hash := hexutil.Encode(hasher.Sum(nil)[:])

	if !strings.HasPrefix(hash, "0x") {
		return "", fmt.Errorf("invalid hex hash: %v", hash)
	}
	if len(hash) < 66 {
		return "", fmt.Errorf("too short length: %v", len(hash))
	}

	return hash[:10], nil
}

func GetConstractParams(function string) ([]*ConstractParams, error) {
	start := strings.Index(function, "(")
	end := strings.Index(function, ")")

	if start >= end {
		return nil, fmt.Errorf("invalid function, start %v >= end %v", start, end)
	}

	params := strings.Split(function[start+1:end], ",")
	constractParams := make([]*ConstractParams, 0)
	for i := range params {
		param := params[i]
		param = strings.TrimSpace(param)
		// todo: strings.Split splits "" to [""]
		if param == "" {
			continue
		}

		tuples := strings.Split(param, " ")
		if len(tuples) == 2 {
			constractParams = append(constractParams, &ConstractParams{Type: tuples[0], Name: tuples[1]})
		} else if len(tuples) == 1 {
			constractParams = append(constractParams, &ConstractParams{Type: tuples[0]})
		} else {
			return nil, fmt.Errorf("invalid params: %v", tuples)
		}
	}

	return constractParams, nil
}

func getFunctionName(function string) string {
	start := strings.Index(function, "(")
	return function[:start]
}

func ConstractMethods() map[string]*RegisteredInputData {
	MethodsRegistry.methods.RLock()
	defer MethodsRegistry.methods.RUnlock()

	return MethodsRegistry.methods.c
}

func SearchConstractMethod(methodID string) (*RegisteredInputData, bool) {
	MethodsRegistry.methods.RLock()
	inputData, ok := MethodsRegistry.methods.c[methodID]
	MethodsRegistry.methods.RUnlock()

	return inputData, ok
}

// GetInputDataByMethodID get InputData by methodID
func GetInputDataByMethodID(methodID string, datas string) (InputData, error) {
	if len(datas)%64 != 0 {
		return nil, fmt.Errorf("invalid length %v of datas", len(datas))
	}

	// search RegisteredInputData firstly
	registeredInputData, ok := SearchConstractMethod(methodID)
	if ok {
		index := 0
		if len(registeredInputData.Params)*64 != len(datas) {
			return nil, fmt.Errorf("datas dont correspond to params, datas %v, params: %v", datas, registeredInputData.Params)
		}

		registeredInputData.MethodID = methodID
		for _, param := range registeredInputData.Params {
			param.Data = fmt.Sprintf("%s%s", "0x", datas[index:index+64])

			index += index + 64
			if index == len(datas) {
				break
			}
		}

		return registeredInputData, nil
	}

	//
	commonInputData := &CommonInputData{MethodID: methodID}
	for i := 0; i < len(datas); i += 64 {
		commonInputData.Params = append(commonInputData.Params, fmt.Sprintf("%s%s", "0x", datas[i:i+64]))
	}

	return commonInputData, nil
}

func HexEncodeByteArray(params []byte) ([]byte, error) {
	buffer := bytes.NewBuffer(params)
	paramsByte, err := cbg.ReadByteArray(buffer, uint64(len(params)))
	if err != nil {
		return nil, err
	}

	hexParams := make([]byte, hex.EncodedLen(len(paramsByte)))
	hex.Encode(hexParams, paramsByte)
	return hexParams, nil
}
