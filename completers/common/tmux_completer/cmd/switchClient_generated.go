package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var switchClientCmd = &cobra.Command{
	Use:   "switch-client",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switchClientCmd).Standalone()

	switchClientCmd.Flags().BoolS("E", "E", false, "TODO description")
	switchClientCmd.Flags().StringS("T", "T", "", "key-table")
	switchClientCmd.Flags().BoolS("Z", "Z", false, "TODO description")
	switchClientCmd.Flags().StringS("c", "c", "", "target-client")
	switchClientCmd.Flags().BoolS("l", "l", false, "TODO description")
	switchClientCmd.Flags().BoolS("n", "n", false, "TODO description")
	switchClientCmd.Flags().BoolS("p", "p", false, "TODO description")
	switchClientCmd.Flags().BoolS("r", "r", false, "TODO description")
	switchClientCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(switchClientCmd)
}
