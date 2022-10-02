package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var unmetCmd = &cobra.Command{
	Use:   "unmet",
	Short: "unmet displays a summary of all unmet dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unmetCmd).Standalone()

	rootCmd.AddCommand(unmetCmd)
}
