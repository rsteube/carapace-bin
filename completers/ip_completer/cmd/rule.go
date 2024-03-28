package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "rule in routing policy database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ruleCmd).Standalone()

	rootCmd.AddCommand(ruleCmd)
}
