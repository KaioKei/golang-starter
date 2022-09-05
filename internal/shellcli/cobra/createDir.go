package cobra

import (
	"fmt"
	"github.com/spf13/cobra"
)

// commands definitions
var createDirCmd = &cobra.Command{
	Use:              "dir",
	Short:            "Create a directory",
	Long:             "Create a directory in the given path. Your user MUST have the proper permissions.",
	TraverseChildren: true, // ensure local flags do not spread to sub commands

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Implement the function to create the directory: %s", Path)
	},
}
