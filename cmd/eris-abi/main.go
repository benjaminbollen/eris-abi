package main

import (
	"fmt"
	"os"

	ebi "github.com/eris-ltd/eris-abi/core"

	log "github.com/eris-ltd/eris-logger"
	. "github.com/eris-ltd/common/go/common"
	"github.com/spf13/cobra"
)

var (
	DefaultInput = "index"
	DefaultIndex = os.Getenv("ERIS_HEAD")

	DefaultHost = "localhost"
	DefaultPort = "4592"
	DefaultAddr = "http://" + DefaultHost + ":" + DefaultPort

	// flags
	IndexFlag string
	InputFlag string
	HostFlag  string
	PortFlag  string
	PPFlag    bool
)

func main() {
	BuildErisABICommand()
	ErisABI.PersistentPreRun = before
	//ErisABI.PersistentPostRun = after
	ErisABI.Execute()
}

var ErisABI = &cobra.Command{
	Use:   "eris-abi",
	Short: "",
	Long:  "",
	Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
}

func BuildErisABICommand() {
	ErisABI.AddCommand(packCmd)
	ErisABI.AddCommand(unpackCmd)
	ErisABI.AddCommand(importCmd)
	ErisABI.AddCommand(addCmd)
	ErisABI.AddCommand(newCmd)
	ErisABI.AddCommand(serverCmd)

	addABIflags()

}

func addABIflags() {
	packCmd.Flags().StringVarP(&InputFlag, "input", "", DefaultInput, "specify input method of ABI data.")
	packCmd.Flags().StringVarP(&IndexFlag, "index", "i", DefaultIndex, "specify index to use as look-up.")

	unpackCmd.Flags().StringVarP(&InputFlag, "input", "", DefaultInput, "specify input method of ABI data.")
	unpackCmd.Flags().StringVarP(&IndexFlag, "index", "i", DefaultIndex, "specify index to use as look-up.")
	unpackCmd.Flags().BoolVarP(&PPFlag, "pp", "p", false, "use json output rather than pretty print.")

	importCmd.Flags().StringVarP(&InputFlag, "input", "", DefaultInput, "specify input method of ABI data.")

	addCmd.Flags().StringVarP(&IndexFlag, "index", "i", DefaultIndex, "specify index to use as look-up.")

	serverCmd.Flags().StringVarP(&PortFlag, "port", "", DefaultPort, "set the port for key daemon to listen on.")
	serverCmd.Flags().StringVarP(&HostFlag, "host", "", DefaultHost, "set the host for key daemon to listen on.")

}

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "generate a transaction",
	Long:  "generate a transaction",
	Run:   cliPack,
}

var unpackCmd = &cobra.Command{
	Use:   "unpack",
	Short: "process contract return values",
	Long:  "process contract return values",
	Run:   cliUnPack,
}

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import an existing ABI file into abi directory",
	Long:  "Import an existing ABI file into abi directory",
	Run:   cliImport,
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an entry to index",
	Long:  "Add an entry to index",
	Run:   cliAdd,
}

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create new index",
	Long:  "Create new index",
	Run:   cliNew,
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts a packing server",
	Long:  "Starts a packing server",
	Run:   cliServer,
}

func before(cmd *cobra.Command, args []string) {

	// log.SetLevel(log.WarnLevel)
	// if do.Verbose {
	// 	log.SetLevel(log.InfoLevel)
	// } else if do.Debug {
	//  log.SetLevel(log.DebugLevel)
	// }

	//Check directory structure exists. If not create it.
	if err := ebi.CheckDirTree(); err != nil {
		//Tree does not exist or is incomplete
		log.Println("Abi directory tree incomplete... Creating it...")
		if err := ebi.BuildDirTree(); err != nil {
			log.Println("Could not build: %s", err)
			IfExit(fmt.Errorf("Could not create directory tree"))
		}
		log.Println("Directory tree built!")
	}

}
