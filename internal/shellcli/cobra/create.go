package cobra

import (
	"github.com/spf13/cobra"
)

// commands definitions
var createCmd = &cobra.Command{
	Use:              "create",                                                                    // command to use
	Short:            "Create a file or a directory",                                              // prompted in parent command helper
	Long:             "Create a file or a directory. Your user MUST have the proper permissions.", // prompted in current command helper
	TraverseChildren: true,                                                                        // ensure local flags do not spread to sub command

	// Run function is mandatory, but can be empty
	Run: func(cmd *cobra.Command, args []string) {},
}
