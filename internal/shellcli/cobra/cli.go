package cobra

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// define here the cli flags
var Verbose bool
var Version bool
var Path string

// rootCmd is the root command definitions
// define here the helper and the root command flags behavior
var rootCmd = &cobra.Command{
	Use:              "cobra",
	Long:             "Shellcli based on cobra",
	TraverseChildren: true, // ensure local flags do not spread to sub commands

	Run: func(cmd *cobra.Command, args []string) {
		// deal with flags
		versionFlag, _ := cmd.Flags().GetBool("version")
		if versionFlag {
			fmt.Println("v0.0.1-alpha")
		}
	},
}

func hello() {
	log.Println("Hello before all !")
}

// used to initialize flags and command
// define the cli architecture here
func init() {
	// can execute functions on cli init
	cobra.OnInitialize(hello)

	// subcommands
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createFileCmd, createDirCmd)

	// optional global flags : accessed by subcommands
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "add more info")             // verbose flag for all commands under root
	createCmd.PersistentFlags().StringVarP(&Path, "path", "p", "", "Path of the resource to create") // path flag for all commands under create

	// local flags : only accessed by root commands
	rootCmd.Flags().BoolVarP(&Version, "version", "V", false, "print the version") // version flag only for root command

	// mark mandatory flags
	err := createCmd.MarkPersistentFlagRequired("path") // mark the flag path as mandatory for all create subcommand
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

// Execute is the entry point of the cli
// You can call it from external packages
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
