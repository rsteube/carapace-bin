package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates and/or modifies files based on a schematic",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()

	generateCmd.Flags().Bool("defaults", false, "Disable interactive input prompts for options with a default.")
	generateCmd.Flags().Bool("dry-run", false, "Run through and reports activity without writing out results.")
	generateCmd.Flags().Bool("force", false, "Force overwriting of existing files.")
	generateCmd.Flags().Bool("interactive", false, "Enable interactive input prompts.")
	rootCmd.AddCommand(generateCmd)
}
