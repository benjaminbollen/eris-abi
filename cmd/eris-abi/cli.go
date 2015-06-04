package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/eris-ltd/eris-abi"
)

func cliPack(c *cli.Context) {
	//Check only one abi specification method has been used
	input := c.String(input)

	args := c.Args()

	
	if (input == "file") {
		//When using file input methos, Get abi from file
		fname := args[0]
		data := args[1:]

		fpath, err := ebi.PathFromHere(fname)
		ifExit(err)

		abiData, _, err := ebi.ReadAbiFile(fpath)
		ifExit(err)

		tx, err := ebi.Packer(abiData, data...)
		ifExit(err)

		fmt.Printf("%s\n", tx)
		return
	} else if (input == "json")
		//When using json input method, read json-abi string from command line
		json := []byte(args[0])
		data := args[1:]

		tx, err := ebi.Packer(json, data...)
		ifExit(err)

		fmt.Printf("%s\n", tx)
		return
	} else if (input == "hash") {
		//Read from the /raw/hash file
		hash := args[0]
		data := args[1:]

		abiData, _, err := ebi.ReadAbiFile(hash)
		ifExit(err)

		tx, err := ebi.Packer(abiData, data...)
		ifExit(err)

		fmt.Printf("%s\n", tx)
		return
	} else if (input == "id") {
		//The id input method uses the indexing system
		hash, err := ebi.IndexResolve(c.String("chainid"), args[0])
		data := args[1:]
		ifExit(err)

		abiData, _, err := ebi.ReadAbiFile(hash)
		ifExit(err)

		tx, err := ebi.Packer(abiData, data...)
		ifExit(err)

		fmt.Printf("%s\n", tx)
		return

	} else {
		err = fmt.Errorf("Unrecognized input method: %s", input)
		ifExit(err)
	}
}