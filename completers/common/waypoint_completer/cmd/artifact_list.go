package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var artifact_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List pushed artifacts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(artifact_listCmd).Standalone()

	artifact_listCmd.Flags().Bool("json", false, "Output the deployment information as JSON.")
	artifact_listCmd.Flags().Bool("long-ids", false, "Show long identifiers rather than sequence numbers.")
	artifact_listCmd.Flags().Bool("verbose", false, "Display more details about each deployment.")
	artifact_listCmd.Flags().Bool("workspace-all", false, "List builds in all workspaces for this project and application.")

	addGlobalOptions(artifact_listCmd)
	addFilterOptions(artifact_listCmd)

	artifactCmd.AddCommand(artifact_listCmd)
}
