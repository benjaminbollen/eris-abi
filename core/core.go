package ebi

import (
	"encoding/hex"
	"os"
	"path"

	"github.com/eris-ltd/eris-abi/abi"

	log "github.com/eris-ltd/eris-abi/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/eris-ltd/eris-abi/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
)

func PathFromHere(fname string) (string, error) {
	//Check for absolute path
	if !path.IsAbs(fname) {
		wd, err := os.Getwd()
		if err != nil {
			return "", err
		}

		return path.Join(wd, fname), nil
	} else {
		return fname, nil
	}
}

//Use the indexing system to pull out file path
func ResolveAbiPath(chainid, contract string) (string, error) {
	return "", nil
}

func MakeAbi(abiData []byte) (abi.ABI, error) {
	if len(abiData) == 0 {
		return abi.NullABI, nil
	}

	abiSpec := new(abi.ABI)
	if err := abiSpec.UnmarshalJSON(abiData); err != nil {
		log.Println("failed to unmarshal", err)
		return abi.NullABI, err
	}

	return *abiSpec, nil
}

//Convenience Packing Functions
func Packer(abiData []byte, data ...string) (string, error) {
	abiSpec, err := MakeAbi(abiData)
	if err != nil {
		return "", err
	}

	tx, err := PackArgsABI(abiSpec, data...)
	if err != nil {
		return "", err
	}

	return tx, nil
}

func PackArgsABI(abiSpec abi.ABI, data ...string) (string, error) {

	funcName := data[0]
	args := data[1:]

	a := []interface{}{}
	for _, aa := range args {
		bb, _ := hex.DecodeString(common.StripHex(common.CoerceHexAndPad(aa, true)))
		a = append(a, bb)
	}

	packedBytes, err := abiSpec.Pack(funcName, args)
	if err != nil {
		return "", err
	}

	packed := hex.EncodeToString(packedBytes)

	return packed, nil
}

func UnPacker(abiData []byte, name string, datas string, pp bool) (string, error) {
	data, _ := hex.DecodeString(datas)

	abiSpec, err := MakeAbi(abiData)
	if err != nil {
		return "", err
	}

	unpacked, err := abiSpec.UnPack(name, data)

	if err != nil {
		return "", err
	}

	if pp {
		return abi.UnpackPrettyPrint(unpacked)
	}
	return string(unpacked), nil
}

// filePack: Read abi data from specified file
func FilePack(filename string, args ...string) (string, error) {
	filepath, err := PathFromHere(filename)
	if err != nil {
		return "", err
	}

	abiData, _, err := ReadAbiFile(filepath)
	if err != nil {
		return "", err
	}

	tx, err := Packer(abiData, args...)
	if err != nil {
		return "", err
	}

	return tx, nil
}

func FileUnPack(filename string, name string, data string, pp bool) (string, error) {
	filepath, err := PathFromHere(filename)
	if err != nil {
		return "", err
	}

	abiData, _, err := ReadAbiFile(filepath)
	if err != nil {
		return "", err
	}

	ups, err := UnPacker(abiData, name, data, pp)
	if err != nil {
		return "", err
	}

	return ups, nil
}

// jsonPack not needed: use Packer

// hashPack: Read abi Data from ebi-tree with supplied hashPack
func HashPack(hash string, args ...string) (string, error) {
	abiData, _, err := ReadAbi(hash)
	if err != nil {
		return "", err
	}

	tx, err := Packer(abiData, args...)
	if err != nil {
		return "", err
	}

	return tx, nil
}

func HashUnPack(hash string, name string, data string, pp bool) (string, error) {
	abiData, _, err := ReadAbi(hash)
	if err != nil {
		return "", err
	}

	ups, err := UnPacker(abiData, name, data, pp)
	if err != nil {
		return "", err
	}

	return ups, nil

}

// indexPack: use the index system to fetch abi data
func IndexPack(index string, key string, args ...string) (string, error) {
	hash, err := IndexResolve(index, key)
	if err != nil {
		return "", err
	}

	abiData, _, err := ReadAbi(hash)
	if err != nil {
		return "", err
	}

	tx, err := Packer(abiData, args...)
	if err != nil {
		return "", err
	}

	return tx, nil
}

func IndexUnPack(index string, key string, name string, data string, pp bool) (string, error) {
	hash, err := IndexResolve(index, key)
	if err != nil {
		return "", err
	}

	abiData, _, err := ReadAbi(hash)
	if err != nil {
		return "", err
	}

	ups, err := UnPacker(abiData, name, data, pp)
	if err != nil {
		return "", err
	}

	return ups, nil
}
