package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var showEnvironmentCmd = &cobra.Command{
	Use:   "show-environment",
	Short: "Dump environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showEnvironmentCmd).Standalone()

	rootCmd.AddCommand(showEnvironmentCmd)
}
