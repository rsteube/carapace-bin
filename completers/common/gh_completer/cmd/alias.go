package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "Create command shortcuts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(aliasCmd).Standalone()
	rootCmd.AddCommand(aliasCmd)
}
