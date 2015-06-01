package main

import (
	"fmt"
	"os"
	"path"

	"github.com/codegangsta/cli"
	"github.com/eris-ltd/epm-go/utils" //Using non-standard directory for abi storage
)

var (
	DefaultDir  = path.Join(utils.Decerver, "contracts")
	DefaultFile = ""

	DefaultHost = "localhost"
	DefaultPort = "4592"
	DefaultAddr = "http://" + DefaultHost + ":" + DefaultPort
)

func main() {
	app := cli.NewApp()
	app.Name = "ebi"
	app.Usage = "Tool for using ABI's to construct transaction data"
	app.Version = "0.0.1"
	app.Author = "Dennis Mckinnon"
	app.Email = "contact@erisindustries.com"
	app.Commands = []cli.Command{
		packCmd,
	}

	app.Run(os.Args)

}

//Excessive structuring to not prohibit future expansion of this tool
var (
	packCmd = cli.Command{
		Name:   "pack",
		Usage:  "generate a transaction",
		Action: cliPack,
		Flags: []cli.Flag{
			fileFlag,
		},
	}

	fileFlag = cli.StringFlag{
		Name:  "file",
		Value: DefaultFile,
		Usage: "Specify the ABI file (Containing JSON ABI) to use",
	}
)

func exit(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func ifExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
