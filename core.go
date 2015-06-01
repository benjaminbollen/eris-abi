package main

import (
	"os"
	"log"
	"fmt"
	"path"
	"strconv"
	"io/ioutil"
	"encoding/hex"
	"github.com/eris-ltd/ebi/abi"
	"github.com/eris-ltd/epm-go/utils"
)

func corePack(fname string, data []string) (string, error){

	fpath, err := ResolveAbiPath(fname)
	if err != nil {
		return "", err
	}

	abiSpec, err := ReadAbi(fpath)
	if err != nil {
		return "", err
	}

	tx, err := packArgsABI(abiSpec, data...)
	if err != nil {
		return "", err
	}

	return tx, nil
}

func ResolveAbiPath(fname string) (string, error){
	//TODO: Handle finding abi stored in eris file structure
	
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Join(wd, fname), nil
}

func ReadAbi(fpath string) (abi.ABI, error) {

	if _, err := os.Stat(fpath); err != nil {
		log.Println("Abi doesn't exist for", fpath)
		return abi.NullABI, err
	}

	abiData, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Println("Failed to read abi file:", err)
		return abi.NullABI, err
	}

	abiSpec := new(abi.ABI)
	if err := abiSpec.UnmarshalJSON(abiData); err != nil {
		log.Println("failed to unmarshal", err)
		return abi.NullABI, err
	}

	return *abiSpec, nil
}

func packArgsABI(abiSpec abi.ABI, data ...string) (string, error) {
//	packed := []string{}

	funcName := data[0]
	args := data[1:]

	a := []interface{}{}
	for _, aa := range args {
		aa = coerceHex(aa, true)
		bb, _ := hex.DecodeString(utils.StripHex(aa))
		a = append(a, bb)
	}

	packedBytes, err := abiSpec.Pack(funcName, a...)
	if err != nil {
		return "", err
	}

	packed := hex.EncodeToString(packedBytes)

	return packed, nil
}

func coerceHex(aa string, padright bool) string {
	if !utils.IsHex(aa) {
		//first try and convert to int
		n, err := strconv.Atoi(aa)
		if err != nil {
			// right pad strings
			if padright {
				aa = "0x" + fmt.Sprintf("%x", aa) + fmt.Sprintf("%0"+strconv.Itoa(64-len(aa)*2)+"s", "")
			} else {
				aa = "0x" + fmt.Sprintf("%x", aa)
			}
		} else {
			aa = "0x" + fmt.Sprintf("%x", n)
		}
	}
	return aa
}