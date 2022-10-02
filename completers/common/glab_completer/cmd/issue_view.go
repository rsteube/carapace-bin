package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var issue_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Display the title, body, and other information about an issue.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(issue_viewCmd).Standalone()
	issue_viewCmd.Flags().BoolP("comments", "c", false, "Show mr comments and activities")
	issue_viewCmd.Flags().IntP("page", "p", 1, "Page number")
	issue_viewCmd.Flags().IntP("per-page", "P", 20, "Number of items to list per page")
	issue_viewCmd.Flags().BoolP("system-logs", "s", false, "Show system activities / logs")
	issue_viewCmd.Flags().BoolP("web", "w", false, "Open mr in a browser. Uses default browser or browser specified in BROWSER variable")
	issueCmd.AddCommand(issue_viewCmd)

	carapace.Gen(issue_viewCmd).PositionalCompletion(
		action.ActionIssues(issue_closeCmd, "opened"),
	)
}
