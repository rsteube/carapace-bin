package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var orgCmd = &cobra.Command{
	Use:   "org",
	Short: "Manage Organization configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(orgCmd).Standalone()
	rootCmd.AddCommand(orgCmd)
}
