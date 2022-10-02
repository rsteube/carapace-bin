package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runner_resourceClass_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a resource-class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runner_resourceClass_createCmd).Standalone()
	runner_resourceClass_createCmd.PersistentFlags().Bool("generate-token", false, "Generate a default token")
	runner_resourceClassCmd.AddCommand(runner_resourceClass_createCmd)
}
