package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var containerCmd = &cobra.Command{
	Use:   "container",
	Short: "Manage containers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(containerCmd).Standalone()

	rootCmd.AddCommand(containerCmd)
}
