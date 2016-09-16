package ebi

import (
	"math/big"
	"fmt"
	"strings"
	"reflect"

	"github.com/eris-ltd/eris-abi/abi"

	log "github.com/eris-ltd/eris-logger"
	"github.com/eris-ltd/common/go/common"
	pmDefinitions "github.com/eris-ltd/eris-pm/definitions"
)

func MakeAbi(abiData string) (abi.ABI, error) {
	if len(abiData) == 0 {
		return abi.ABI{}, nil
	}

	abiSpec, err := abi.JSON(strings.NewReader(abiData))
	if err != nil {
		return abi.ABI{}, err
	}

	return abiSpec, nil
}

//Convenience Packing Functions
func Packer(abiData, funcName string, args ...interface{}) ([]byte, error) {
	abiSpec, err := MakeAbi(abiData)
	if err != nil {
		return nil, err
	}

	packedBytes, err := abiSpec.Pack(funcName, args)
	if err != nil {
		return nil, err
	}

	return packedBytes, nil
}

func Unpacker(abiData, name string, data []byte) ([]pmDefinitions.Variable, error) {

	abiSpec, err := MakeAbi(abiData)
	if err != nil {
		return []pmDefinitions.Variable{}, err
	}

	//unpacked, err := createOutputInterface(abiSpec, name)
	//if err != nil {
	//	return nil, err
	//}
	var unpacked interface{}
	err = abiSpec.Unpack(&unpacked, name, data)
	if err != nil {
		return []pmDefinitions.Variable{}, err
	}

	return formatUnpackedReturn(abiSpec, name, unpacked)
}

func createOutputInterface(abiSpec abi.ABI, methodName string) ([]interface{}, error) {
	var outputInterface []interface{}

	method, exist := abiSpec.Methods[methodName]	
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", methodName)
	}
	
	if len(method.Outputs) == 0 {
		log.Debug("Empty output, nothing to interface to")
		return nil, nil
	}
	
	for _, output := range method.Outputs {
		typ := output.Type
		outputInterface = append(outputInterface, typ)
	}
	return outputInterface, nil
}

func formatUnpackedReturn(abiSpec abi.ABI, methodName string, values ...interface{}) ([]pmDefinitions.Variable, error) {
	var returnVars []pmDefinitions.Variable
	method, exist := abiSpec.Methods[methodName]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", methodName)
	}
	for i, output := range method.Outputs {
		v := reflect.ValueOf(values[i])

		var StringVal string
		fmt.Println("Kind: ", v.Kind().String())
		fmt.Println(values[i])
		if v.Kind() == reflect.String {
			StringVal = fmt.Sprintf("%v", values[i])
		} else if v.Kind() == reflect.Ptr {
			bigInt := v.Interface().(*big.Int)
			switch output.Type.String()[:3] {
			case "int":
				StringVal = common.S256(bigInt).String()
			case "uin":
				StringVal = common.U256(bigInt).String()
			}	
		} else {
			StringVal = fmt.Sprintf("%v", values[i])
		}
		name := output.Name
		if output.Name == "" {
			name = string(i)
		}
		arg := pmDefinitions.Variable{
				Name: name,
				Value: StringVal,
			}
		returnVars = append(returnVars, arg)
	}
	return returnVars, nil
}