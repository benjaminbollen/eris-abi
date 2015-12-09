package main

import (
	"fmt"

	ebi "github.com/eris-ltd/eris-abi/core"

	"github.com/eris-ltd/eris-abi/Godeps/_workspace/src/github.com/codegangsta/cli"
)

func cliPack(c *cli.Context) {
	input := c.String("input")

	args := c.Args()

	if input == "file" {
		//When using file input methos, Get abi from file
		fname := args[0]
		data := args[1:]

		tx, err := ebi.FilePack(fname, data...)
		ifExit(err)

		logger.Printf("%s\n", tx)
		return
	} else if input == "json" {
		//When using json input method, read json-abi string from command line
		json := []byte(args[0])
		data := args[1:]

		tx, err := ebi.Packer(json, data...)
		ifExit(err)

		logger.Printf("%s\n", tx)
		return
	} else if input == "hash" {
		//Read from the /raw/hash file
		hash := args[0]
		data := args[1:]

		tx, err := ebi.HashPack(hash, data...)
		ifExit(err)

		logger.Printf("%s\n", tx)
		return
	} else if input == "index" {
		//The index input method uses the indexing system
		index := c.String("index")
		key := args[0]
		data := args[1:]

		tx, err := ebi.IndexPack(index, key, data...)
		ifExit(err)

		logger.Printf("%s\n", tx)
		return

	} else {
		err := fmt.Errorf("Unrecognized input method: %s\n", input)
		ifExit(err)
	}
}

func cliUnPack(c *cli.Context) {
	input := c.String("input")
	pp := c.Bool("pp")
	args := c.Args()

	if input == "file" {
		//When using file input methods, Get abi from file
		fname := args[0]
		name := args[1]
		data := args[2]

		abs, err := ebi.FileUnPack(fname, name, data, pp)
		ifExit(err)

		logger.Printf("%s\n", abs)
		return
	} else if input == "json" {
		//When using json input method, read json-abi string from command line
		json := []byte(args[0])
		name := args[1]
		data := args[2]

		abs, err := ebi.UnPacker(json, name, data, pp)
		ifExit(err)

		logger.Printf("%s\n", abs)
		return
	} else if input == "hash" {
		//Read from the /raw/hash file
		hash := args[0]
		name := args[1]
		data := args[2]

		abs, err := ebi.HashUnPack(hash, name, data, pp)
		ifExit(err)

		logger.Printf("%s\n", abs)
		return
	} else if input == "index" {
		//The index input method uses the indexing system
		index := c.String("index")
		key := args[0]
		name := args[1]
		data := args[2]

		abs, err := ebi.IndexUnPack(index, key, name, data, pp)
		ifExit(err)

		logger.Printf("%s\n", abs)
		return

	} else {
		err := fmt.Errorf("Unrecognized input method: %s\n", input)
		ifExit(err)
	}
}

func cliImport(c *cli.Context) {
	//Import an abi file into abi directory
	args := c.Args()

	if c.String("input") == "file" {
		fname := args[0]

		fpath, err := ebi.PathFromHere(fname)
		ifExit(err)

		abiData, abiHash, err := ebi.ReadAbiFile(fpath)
		ifExit(err)

		_, err = ebi.WriteAbi(abiData)
		ifExit(err)

<<<<<<< 59ea82c6982e0028840d52b6b3a7a055c7ca6aee
		fmt.Printf("Imported Abi as %s\n", abiHash)
=======
		logger.Printf("Imported Abi as %s\n", abiHash)
>>>>>>> ...
	} else if c.String("input") == "json" {
		json := []byte(args[0])
		abiHash, err := ebi.WriteAbi(json)
		ifExit(err)

		logger.Printf("Imported Abi as %s\n", abiHash)
	}
	return
}

func cliAdd(c *cli.Context) {
	//Add an entry to index
	args := c.Args()
	iname := c.String("index")
	key := args[0]
	value := args[1]

	err := ebi.AddEntry(iname, key, value)
	ifExit(err)

	logger.Printf("Added Entry %s as %s\n", value, key)
	return
}

func cliNew(c *cli.Context) {
	//Create new index
	args := c.Args()
	iname := args[0]

	err := ebi.NewIndex(iname)
	ifExit(err)

	logger.Printf("Created new index: %s\n", iname)
	return
}

func cliServer(c *cli.Context) {
	host, port := c.String("host"), c.String("port")
	ifExit(ListenAndServe(host, port))
}
