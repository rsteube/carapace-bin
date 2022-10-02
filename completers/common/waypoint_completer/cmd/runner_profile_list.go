package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_profile_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all registered runner profiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_profile_listCmd).Standalone()

	runner_profileCmd.AddCommand(runner_profile_listCmd)
}
