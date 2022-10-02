package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var artifact_pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push a build's artifact to a registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_pushCmd).Standalone()

	addGlobalOptions(artifact_pushCmd)
	addOperationOptions(artifact_pushCmd)

	artifactCmd.AddCommand(artifact_pushCmd)
}
