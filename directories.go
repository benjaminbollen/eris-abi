package ebi

import (
	"os"
	"fmt"
	"path"
	"github.com/eris-ltd/eris-cli/util"
)

func GetAbiRoot() string {
	var abiroot string
	if (os.Getenv("ERIS_ABI_ROOT") != "") {
		abiroot = os.Getenv("ERIS_ABI_ROOT")
	} else {
		path.Join(util.UserErisDir(), "abi")
	}
	return abiroot
}

//Directory Structure
//.eris
//   +-abi
//      +-index
//      |    + abi indexing jsons
//      |
//      +-raw
//           + abi files (hash named)

var (
	Root = GetAbiRoot()
	Index = path.Join(Root, "index")
	Raw = path.Join(Root, "raw")
)

func BuildDirTree() {
	//Check if abi root exists.
	if _, err := os.Stat(Root); err != nil {
		//create it
		err = os.MkdirAll(Root, 0700)
		if err != nil {
			fmt.Errorf("Failed to create the abi root directory")
		}
	}

	//Create Indexing folder
	if _, err := os.Stat(Index); err != nil {
		//create it
		err = os.MkdirAll(Index, 0700)
		if err != nil {
			fmt.Errorf("Failed to create the abi root directory")
		}
	}

	//Create Raw Storage Folder
	if _, err := os.Stat(Raw); err != nil {
		//create it
		err = os.MkdirAll(Raw, 0700)
		if err != nil {
			fmt.Errorf("Failed to create the abi root directory")
		}
	}
}