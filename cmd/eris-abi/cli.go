package main

import (
	"fmt"

	ebi "github.com/eris-ltd/eris-abi/core"

	log "github.com/eris-ltd/eris-abi/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	. "github.com/eris-ltd/eris-abi/Godeps/_workspace/src/github.com/eris-ltd/common/go/common"
	"github.com/eris-ltd/eris-abi/Godeps/_workspace/src/github.com/spf13/cobra"
)

func cliPack(cmd *cobra.Command, args []string) {
	input := InputFlag

	if input == "file" {
		//When using file input methos, Get abi from file
		fname := args[0]
		data := args[1:]

		tx, err := ebi.FilePack(fname, data...)
		IfExit(err)

		log.Printf("%s\n", tx)
		return
	} else if input == "json" {
		//When using json input method, read json-abi string from command line
		json := []byte(args[0])
		data := args[1:]

		tx, err := ebi.Packer(json, data...)
		IfExit(err)

		log.Printf("%s\n", tx)
		return
	} else if input == "hash" {
		//Read from the /raw/hash file
		hash := args[0]
		data := args[1:]

		tx, err := ebi.HashPack(hash, data...)
		IfExit(err)

		log.Printf("%s\n", tx)
		return
	} else if input == "index" {
		//The index input method uses the indexing system
		index := IndexFlag
		key := args[0]
		data := args[1:]

		tx, err := ebi.IndexPack(index, key, data...)
		IfExit(err)

		log.Printf("%s\n", tx)
		return

	} else {
		err := fmt.Errorf("Unrecognized input method: %s\n", input)
		IfExit(err)
	}
}

func cliUnPack(cmd *cobra.Command, args []string) {
	input := InputFlag
	pp := PPFlag

	if input == "file" {
		//When using file input methods, Get abi from file
		fname := args[0]
		name := args[1]
		data := args[2]

		abs, err := ebi.FileUnPack(fname, name, data, pp)
		IfExit(err)

		log.Printf("%s\n", abs)
		return
	} else if input == "json" {
		//When using json input method, read json-abi string from command line
		json := []byte(args[0])
		name := args[1]
		data := args[2]

		abs, err := ebi.UnPacker(json, name, data, pp)
		IfExit(err)

		log.Printf("%s\n", abs)
		return
	} else if input == "hash" {
		//Read from the /raw/hash file
		hash := args[0]
		name := args[1]
		data := args[2]

		abs, err := ebi.HashUnPack(hash, name, data, pp)
		IfExit(err)

		log.Printf("%s\n", abs)
		return
	} else if input == "index" {
		//The index input method uses the indexing system
		index := IndexFlag
		key := args[0]
		name := args[1]
		data := args[2]

		abs, err := ebi.IndexUnPack(index, key, name, data, pp)
		IfExit(err)

		log.Printf("%s\n", abs)
		return

	} else {
		err := fmt.Errorf("Unrecognized input method: %s\n", input)
		IfExit(err)
	}
}

func cliImport(cmd *cobra.Command, args []string) {
	//Import an abi file into abi directory

	if InputFlag == "file" {
		fname := args[0]

		fpath, err := ebi.PathFromHere(fname)
		IfExit(err)

		abiData, abiHash, err := ebi.ReadAbiFile(fpath)
		IfExit(err)

		_, err = ebi.WriteAbi(abiData)
		IfExit(err)

		log.Printf("Imported Abi as %s\n", abiHash)
	} else if InputFlag == "json" {
		json := []byte(args[0])
		abiHash, err := ebi.WriteAbi(json)
		IfExit(err)

		log.Printf("Imported Abi as %s\n", abiHash)
	}
	return
}

func cliAdd(cmd *cobra.Command, args []string) {
	//Add an entry to index
	iname := IndexFlag
	key := args[0]
	value := args[1]

	err := ebi.AddEntry(iname, key, value)
	IfExit(err)

	log.Printf("Added Entry %s as %s\n", value, key)
	return
}

func cliNew(cmd *cobra.Command, args []string) {
	//Create new index
	iname := args[0]

	err := ebi.NewIndex(iname)
	IfExit(err)

	log.Printf("Created new index: %s\n", iname)
	return
}

func cliServer(cmd *cobra.Command, args []string) {
	host, port := HostFlag, PortFlag
	IfExit(ListenAndServe(host, port))
}
