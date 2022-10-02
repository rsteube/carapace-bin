package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize and validate a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().String("from-project", "", "Create a new application by fetching the given application from a remote source.")
	initCmd.Flags().String("into", "", "Where to write the application fetched via -from-project")
	initCmd.Flags().Bool("update", false, "Update the project configuration if it already exists.")

	addGlobalOptions(initCmd)

	rootCmd.AddCommand(initCmd)
}
