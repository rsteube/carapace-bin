package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var pr_checksCmd = &cobra.Command{
	Use:   "checks",
	Short: "Show CI status for a single pull request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_checksCmd).Standalone()
	pr_checksCmd.Flags().IntP("interval", "i", 10, "Refresh interval in seconds when using `--watch` flag")
	pr_checksCmd.Flags().Bool("required", false, "Only show checks that are required")
	pr_checksCmd.Flags().Bool("watch", false, "Watch checks until they finish")
	pr_checksCmd.Flags().BoolP("web", "w", false, "Open the web browser to show details about checks")
	prCmd.AddCommand(pr_checksCmd)

	carapace.Gen(pr_checksCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				action.ActionPullRequests(pr_checksCmd, action.PullRequestOpts{Open: true}),
				action.ActionBranches(pr_checksCmd),
			).Invoke(c).Merge().ToA()
		}),
	)
}
