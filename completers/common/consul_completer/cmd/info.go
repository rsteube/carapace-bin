package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Provides debugging information for operators",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(infoCmd).Standalone()
	addClientFlags(infoCmd)

	rootCmd.AddCommand(infoCmd)
}
