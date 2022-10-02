package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var ci_lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Checks if your .gitlab-ci.yml file is valid.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ci_lintCmd).Standalone()
	ciCmd.AddCommand(ci_lintCmd)

	carapace.Gen(ci_lintCmd).PositionalCompletion(
		carapace.ActionFiles(".gitlab-ci.yml"),
	)
}
