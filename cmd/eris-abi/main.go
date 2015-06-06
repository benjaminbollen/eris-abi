package main

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
	"github.com/eris-ltd/eris-abi"
)

var (
	DefaultInput = "index"
	DefaultIndex = os.Getenv("ERIS_HEAD") 

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
		importCmd,
		addCmd,
		newCmd,
		serverCmd,
	}

	app.Before = func(c *cli.Context) error{
		//Check directory structure exists. If not create it.
		err := ebi.CheckDirTree()
		if err != nil {
			//Tree does not exist or is incomplete
			fmt.Println("Abi directory tree incomplete... Creating it...")
			err := ebi.BuildDirTree()
			if err != nil {
				fmt.Println("Could not build: %s", err)
				return fmt.Errorf("Could not create directory tree")
			}
			fmt.Println("Directory tree built!")
		}

		return nil
	}

	app.Run(os.Args)

}

var (
	packCmd = cli.Command{
		Name:   "pack",
		Usage:  "generate a transaction",
		Action: cliPack,
		Flags: []cli.Flag{
			inputFlag,
			indexFlag,
		},
	}

	importCmd = cli.Command{
		Name:	"import",
		Usage:	"Import an existing ABI file into abi directory",
		Action:	cliImport,
	}

	addCmd = cli.Command{
		Name: 	"add",
		Usage:	"Add an entry to index",
		Action:	cliAdd,
	}

	newCmd = cli.Command{
		Name: 	"new",
		Usage:	"Create new index",
		Action:	cliNew,
	}

	serverCmd = cli.Command{
		Name:	"server",
		Usage:	"Starts a packing server",
		Action:	cliServer,
		Flags:	[]cli.Flag{
			hostFlag,
			portFlag,
		},
	}

	inputFlag = cli.StringFlag{
		Name: "input",
		Value: DefaultInput,
		Usage: "Specify input method of ABI data.",
	}

	indexFlag = cli.StringFlag{
		Name: 	"index, i",
		Usage:	"Specify Chainid to use as look-up",
		Value:	DefaultIndex,
	}

	portFlag = cli.StringFlag{
		Name:  "port",
		Value: DefaultPort,
		Usage: "set the port for key daemon to listen on",
	}

	hostFlag = cli.StringFlag{
		Name:  "host",
		Value: DefaultHost,
		Usage: "set the host for key daemon to listen on",
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
