package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mr_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display the title, body, and other information about a merge request.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mr_viewCmd).Standalone()
	mr_viewCmd.Flags().BoolP("comments", "c", false, "Show mr comments and activities")
	mr_viewCmd.Flags().IntP("page", "p", 0, "Page number")
	mr_viewCmd.Flags().IntP("per-page", "P", 20, "Number of items to list per page")
	mr_viewCmd.Flags().BoolP("system-logs", "s", false, "Show system activities / logs")
	mr_viewCmd.Flags().BoolP("web", "w", false, "Open mr in a browser. Uses default browser or browser specified in BROWSER variable")
	mrCmd.AddCommand(mr_viewCmd)

	carapace.Gen(mr_viewCmd).PositionalCompletion(
		action.ActionMergeRequestsAndBranches(mr_viewCmd, "opened"),
	)
}
