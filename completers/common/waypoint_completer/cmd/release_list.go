package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var release_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List releases",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_listCmd).Standalone()

	release_listCmd.Flags().Bool("json", false, "Output the release information as JSON.")
	release_listCmd.Flags().Bool("long-ids", false, "Show long identifiers rather than sequence numbers.")
	release_listCmd.Flags().Bool("url", false, "Display release URL.")
	release_listCmd.Flags().Bool("verbose", false, "Display more details about each release.")
	release_listCmd.Flags().Bool("workspace-all", false, "List builds in all workspaces for this project and application.")

	addFilterOptions(release_listCmd)

	releaseCmd.AddCommand(release_listCmd)
}
