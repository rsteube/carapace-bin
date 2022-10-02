package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var artifact_listBuildsCmd = &cobra.Command{
	Use:   "list-builds",
	Short: "List builds",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_listBuildsCmd).Standalone()

	artifact_listBuildsCmd.Flags().Bool("long-ids", false, "Show long identifiers rather than sequence numbers.")
	artifact_listBuildsCmd.Flags().Bool("workspace-all", false, "List builds in all workspaces for this project and application.")

	addGlobalOptions(artifact_listBuildsCmd)

	artifactCmd.AddCommand(artifact_listBuildsCmd)
}
