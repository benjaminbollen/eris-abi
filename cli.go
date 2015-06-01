package main

import (
	"fmt"
	"github.com/codegangsta/cli"
)

func cliPack(c *cli.Context) {
	file := c.String("file")

	tx, err := corePack(file, c.Args())
	ifExit(err)
	fmt.Printf("%s\n", tx)
}