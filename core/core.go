package ebi

import (
	"fmt"
	"strings"
	"reflect"

	"github.com/eris-ltd/eris-abi/abi"

	log "github.com/eris-ltd/eris-logger"
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

func Unpacker(abiData, name string, data []byte) ([]*pmDefinitions.Variable, error) {

	abiSpec, err := MakeAbi(abiData)
	if err != nil {
		return []*pmDefinitions.Variable{}, err
	}

	unpacked, err := createOutputInterface(abiSpec, name)
	err = abiSpec.Unpack(&unpacked, name, data)
	if err != nil {
		return []*pmDefinitions.Variable{}, err
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
		typ := output.Type.Type
		outputInterface = append(outputInterface, reflect.New(typ))
	}
	return outputInterface, nil
}

func formatUnpackedReturn(abiSpec abi.ABI, methodName string, values []interface{}) ([]*pmDefinitions.Variable, error) {
	var returnVars []*pmDefinitions.Variable
	method, exist := abiSpec.Methods[methodName]
	if !exist {
		return nil, fmt.Errorf("method '%s' not found", methodName)
	}
	for i, output := range method.Outputs {
		arg := &pmDefinitions.Variable{
				Name: output.Name,
				Value: fmt.Sprintf("%v", values[i]),
			}
		returnVars = append(returnVars, arg)
	}
	return returnVars, nil
}