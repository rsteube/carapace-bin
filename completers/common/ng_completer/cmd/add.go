package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds support for an external library to your project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().Bool("defaults", false, "Disable interactive input prompts for options with a default.")
	addCmd.Flags().Bool("interactive", false, "Enable interactive input prompts.")
	addCmd.Flags().String("registry", "", "The NPM registry to use.")
	addCmd.Flags().Bool("skip-confirmation", false, "Skip asking a confirmation prompt before installing and executing the package.")
	addCmd.Flags().Bool("verbose", false, "Display additional details about internal operations during execution.")
	rootCmd.AddCommand(addCmd)

	// TODO positional completion
}
