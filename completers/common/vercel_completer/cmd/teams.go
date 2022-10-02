package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var teamsCmd = &cobra.Command{
	Use:   "teams",
	Short: "Manages your teams",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(teamsCmd).Standalone()

	rootCmd.AddCommand(teamsCmd)
}
