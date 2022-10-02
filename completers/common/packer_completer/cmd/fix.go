package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var fixCmd = &cobra.Command{
	Use:   "fix",
	Short: "fixes templates from old versions of packer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fixCmd).Standalone()

	fixCmd.Flags().BoolS("validate", "validate", false, "If true (default), validates the fixed template.")
	rootCmd.AddCommand(fixCmd)

	carapace.Gen(fixCmd).PositionalCompletion(
		carapace.ActionFiles(".json", ".hcl"),
	)
}
