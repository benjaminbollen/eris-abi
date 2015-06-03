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

		fpath, err := ebi.ResolveAbiPath(fname)
		if err != nil {
			return
		}

		abiSpec, err := ebi.ReadFileAbi(fpath)
		if err != nil {
			return
		}
		tx, err := ebi.PackArgsABI(abiSpec, args...)
		ifExit(err)
		fmt.Printf("%s\n", tx)
		return
	}

	//When using --json flag, directly feed json into abi conversion
	if (c.IsSet("json")) {
		json := c.String("json")
		fmt.Println(json)

		abiSpec, err := ebi.ReadJsonAbi(json)
		if err != nil {
			return
		}
		tx, _ := ebi.PackArgsABI(abiSpec, args...)

		ifExit(err)
		fmt.Printf("%s\n", tx)
		return
	}
}