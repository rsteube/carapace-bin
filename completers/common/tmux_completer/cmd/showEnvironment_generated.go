package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var showEnvironmentCmd = &cobra.Command{
	Use:   "show-environment",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showEnvironmentCmd).Standalone()

	showEnvironmentCmd.Flags().BoolS("g", "g", false, "TODO description")
	showEnvironmentCmd.Flags().BoolS("h", "h", false, "TODO description")
	showEnvironmentCmd.Flags().BoolS("s", "s", false, "TODO description")
	showEnvironmentCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(showEnvironmentCmd)
}
