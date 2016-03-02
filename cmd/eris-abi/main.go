package main

import (
	"fmt"
	"os"

	ebi "github.com/eris-ltd/eris-abi/core"

	log "github.com/eris-ltd/eris-abi/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	"github.com/eris-ltd/eris-abi/Godeps/_workspace/src/github.com/codegangsta/cli"
	logger "github.com/eris-ltd/eris-abi/Godeps/_workspace/src/github.com/eris-ltd/common/go/log"
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
		unpackCmd,
		importCmd,
		addCmd,
		newCmd,
		serverCmd,
	}

	app.Before = func(c *cli.Context) error {
		log.SetFormatter(logger.ErisFormatter{})

		// log.SetLevel(log.WarnLevel)
		// if do.Verbose {
		// 	log.SetLevel(log.InfoLevel)
		// } else if do.Debug {
		//  log.SetLevel(log.DebugLevel)
		// }

		//Check directory structure exists. If not create it.
		err := ebi.CheckDirTree()
		if err != nil {
			//Tree does not exist or is incomplete
			log.Println("Abi directory tree incomplete... Creating it...")
			err := ebi.BuildDirTree()
			if err != nil {
				log.Println("Could not build: %s", err)
				return fmt.Errorf("Could not create directory tree")
			}
			log.Println("Directory tree built!")
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

	unpackCmd = cli.Command{
		Name:   "unpack",
		Usage:  "process contract return values",
		Action: cliUnPack,
		Flags: []cli.Flag{
			inputFlag,
			indexFlag,
			ppFlag,
		},
	}

	importCmd = cli.Command{
		Name:   "import",
		Usage:  "Import an existing ABI file into abi directory",
		Action: cliImport,
		Flags: []cli.Flag{
			inputFlag,
		},
	}

	addCmd = cli.Command{
		Name:   "add",
		Usage:  "Add an entry to index",
		Action: cliAdd,
		Flags: []cli.Flag{
			indexFlag,
		},
	}

	newCmd = cli.Command{
		Name:   "new",
		Usage:  "Create new index",
		Action: cliNew,
	}

	serverCmd = cli.Command{
		Name:   "server",
		Usage:  "Starts a packing server",
		Action: cliServer,
		Flags: []cli.Flag{
			hostFlag,
			portFlag,
		},
	}

	inputFlag = cli.StringFlag{
		Name:  "input",
		Value: DefaultInput,
		Usage: "Specify input method of ABI data.",
	}

	indexFlag = cli.StringFlag{
		Name:  "index, i",
		Usage: "Specify index to use as look-up",
		Value: DefaultIndex,
	}

	ppFlag = cli.BoolTFlag{
		Name:  "pp, p",
		Usage: "Turn off pretty print and use json output instead",
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
	log.Println(err)
	os.Exit(1)
}

func ifExit(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
