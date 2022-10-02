package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var builderCmd = &cobra.Command{
	Use:   "builder",
	Short: "Manage builds",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(builderCmd).Standalone()

	rootCmd.AddCommand(builderCmd)
}
