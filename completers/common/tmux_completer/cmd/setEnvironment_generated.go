package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setEnvironmentCmd = &cobra.Command{
	Use:   "set-environment",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setEnvironmentCmd).Standalone()

	setEnvironmentCmd.Flags().BoolS("F", "F", false, "TODO description")
	setEnvironmentCmd.Flags().BoolS("g", "g", false, "TODO description")
	setEnvironmentCmd.Flags().BoolS("h", "h", false, "TODO description")
	setEnvironmentCmd.Flags().BoolS("r", "r", false, "TODO description")
	setEnvironmentCmd.Flags().StringS("t", "t", "", "target-session")
	setEnvironmentCmd.Flags().BoolS("u", "u", false, "TODO description")
	rootCmd.AddCommand(setEnvironmentCmd)
}
