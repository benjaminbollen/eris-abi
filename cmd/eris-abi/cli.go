package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/eris-ltd/eris-abi"
)

func cliPack(c *cli.Context) {
	input := c.String("input")

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
	} else if (input == "json") {
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

		abiData, _, err := ebi.ReadAbi(hash)
		ifExit(err)

		tx, err := ebi.Packer(abiData, data...)
		ifExit(err)

		fmt.Printf("%s\n", tx)
		return
	} else if (input == "index") {
		//The id input method uses the indexing system
		hash, err := ebi.IndexResolve(c.String("index"), args[0])
		data := args[1:]
		ifExit(err)

		abiData, _, err := ebi.ReadAbi(hash)
		ifExit(err)

		tx, err := ebi.Packer(abiData, data...)
		ifExit(err)

		fmt.Printf("%s\n", tx)
		return

	} else {
		err := fmt.Errorf("Unrecognized input method: %s\n", input)
		ifExit(err)
	}
}

func cliImport(c *cli.Context) {
	//Import an abi file into abi directory
	args := c.Args()
	fname := args[0]

	fpath, err := ebi.PathFromHere(fname)
	ifExit(err)

	abiData, abiHash, err := ebi.ReadAbiFile(fpath)
	ifExit(err)

	_, err = ebi.WriteAbi(abiData)
	ifExit(err)

	fmt.Printf("Imported Abi as %s\n", abiHash)
	return
}

func cliAdd(c *cli.Context) {
	//Add an entry to index
	args := c.Args()
	iname := args[0]
	key := args[1]
	value := args[2]

	err := ebi.AddEntry(iname, key, value)
	ifExit(err)

	fmt.Printf("Added Entry %s as %s\n", value, key)
	return
}

func cliNew(c *cli.Context) {
	//Create new index
	args := c.Args()
	iname := args[0]

	err := ebi.NewIndex(iname)
	ifExit(err)

	fmt.Printf("Created new index: %s\n", iname)
	return
}

func cliServer(c *cli.Context) {
	host, port := c.String("host"), c.String("port")
	ifExit(ListenAndServe(host, port))
}