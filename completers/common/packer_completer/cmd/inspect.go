package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "see components of a template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspectCmd).Standalone()

	inspectCmd.Flags().BoolS("machine-readable", "machine-readable", false, "Machine-readable output")
	rootCmd.AddCommand(inspectCmd)

	carapace.Gen(inspectCmd).PositionalCompletion(
		carapace.ActionFiles(".json", ".hcl"),
	)
}
