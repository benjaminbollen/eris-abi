package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/eris-ltd/eris-abi"
)

func cliPack(c *cli.Context) {
	//Check only one abi specification method has been used
	if (c.IsSet("file") && c.IsSet("json")) {
		fmt.Println("ERROR: Can not simultaneously use --json and --file")
		return
	}

	args := c.Args()

	//When using --file flag, Get abi from file
	if (c.IsSet("file")) {
		fname := c.String("file")

		fpath, err := ebi.PathFromHere(fname)
		ifExit(err)

		abiData, _, err := ebi.ReadAbiFile(fpath)
		ifExit(err)

		tx, err := ebi.Packer(abiData, args...)
		ifExit(err)

		fmt.Printf("%s\n", tx)
		return
	}

	//When using --json flag, directly feed json into abi conversion
	if (c.IsSet("json")) {
		json := []byte(c.String("json"))

		tx, err := ebi.Packer(json, args...)
		ifExit(err)

		fmt.Printf("%s\n", tx)
		return
	}
}