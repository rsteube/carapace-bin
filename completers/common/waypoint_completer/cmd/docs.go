package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Show documentation for components",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsCmd).Standalone()

	docsCmd.Flags().Bool("builtin", false, "Show documentation on all builtin plugins The default is false.")
	docsCmd.Flags().Bool("markdown", false, "Show documentation in markdown format The default is false.")
	docsCmd.Flags().String("plugin", "", "Only show documentation for plugins with this name")
	docsCmd.Flags().String("type", "", "Only show documentation for this type of plugin")

	addGlobalOptions(docsCmd)
	addOperationOptions(docsCmd)

	rootCmd.AddCommand(docsCmd)
}
