package ebi

import (
	"os"
	"fmt"
	"path"
	"io/ioutil"
	"encoding/hex"
	"crypto/sha256"
	"github.com/eris-ltd/eris-cli/util"
)

//This file is for storage and retrieval functions of abi's in abi subdirectory

//Write ABI []byte data into hash-named file 
func WriteAbiFile(abiData []byte) error {
	//Construct file path based on data hash
	abiHash := hex.EncodeToString(sha256.Sum(abiData))
	abiPath := path.Join(Raw, abiHash)

	//Write data
	err := ioutil.WriteFile(abiPath, abiData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func ReadAbiFile(abiPath string) ([]byte, error) {
	//Check it exists first
	if _, err := os.Stat(abiPath); err != nil {
		return nil, fmt.Errorf("Could not read ABI file %s: Does not Exist", abiPath)
	}

	abiData, err := ioutil.ReadFile(abiPath)
	if err != nil {
		log.Println("Failed to read abi file:", err)
		return nil, err
	}

	return abiData, nil
}

func ReadAbiHash(abiHash string) ([]byte, error) {
	abiPath := path.Join(Raw, abiHash)

	abiData, err := ReadAbiFile(abiPath)
	if err != nil {
		return nil, err
	}

	return abiData, nil
}

func VerifyAbiHash(abiPath, abiHash string) error {
	abiData, err := ReadAbiFile(abiPath)
	hash := hex.EncodeToString(sha256.Sum(abiData))

	if (hash != abiHash) {
		return fmt.Errorf("The abi data does not match its hash")
	}

	return nil
}

func VerifyAbiFile(abiPath string) error {
	//Get the filename for hash
	finfo, err := os.Stat(abiPath)
	if err != nil {
		return fmt.Errorf("Could not read ABI file %s: Does not Exist", abiPath)
	}

	filename = finfo.Name()

	//Check filename is a hexstring
	_, err := hex.DecodeString(filename)
	if err != nil {
		return fmt.Errorf("File name %s is not a hex string. Can't Compare" filename)
	}

	return VerifyAbiHash(abiPath, filename)
}