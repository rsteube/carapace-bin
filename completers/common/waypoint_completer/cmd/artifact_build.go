package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var artifact_buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build a new versioned artifact from source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_buildCmd).Standalone()

	artifact_buildCmd.Flags().Bool("push", false, "Push the artifact to the configured registry.")

	addGlobalOptions(artifact_buildCmd)
	addOperationOptions(artifact_buildCmd)

	artifactCmd.AddCommand(artifact_buildCmd)
}
