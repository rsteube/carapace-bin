package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var env_pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Pull all Development Environment Variables from the cloud and write to a file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(env_pullCmd).Standalone()

	envCmd.AddCommand(env_pullCmd)

	carapace.Gen(env_pullCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
