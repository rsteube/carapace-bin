package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var override_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List directory toolchain overrides",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(override_listCmd).Standalone()

	override_listCmd.Flags().BoolP("help", "h", false, "Prints help information")
	overrideCmd.AddCommand(override_listCmd)
}
