package main

import (
	"fmt"
	"github.com/codegangsta/cli"
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

		fpath, err := ResolveAbiPath(fname)
		if err != nil {
			return
		}

		abiSpec, err := ReadFileAbi(fpath)
		if err != nil {
			return
		}
		tx, err := packArgsABI(abiSpec, args...)
		ifExit(err)
		fmt.Printf("%s\n", tx)
		return
	}

	//When using --json flag, directly feed json into abi conversion
	if (c.IsSet("json")) {
		json := c.String("json")
		fmt.Println(json)

		abiSpec, err := ReadJsonAbi(json)
		if err != nil {
			return
		}
		tx, _ := packArgsABI(abiSpec, args...)

		ifExit(err)
		fmt.Printf("%s\n", tx)
		return
	}
}