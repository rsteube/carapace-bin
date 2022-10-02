package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the tool to the latest version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()
	updateCmd.PersistentFlags().Bool("check", false, "Check if there are any updates available without installing")
	rootCmd.AddCommand(updateCmd)
}
