package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "Fetch from and integrate with another repository or a local branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pullCmd).Standalone()

	pullCmd.Flags().Bool("all", false, "fetch from all remotes")
	pullCmd.Flags().Bool("allow-unrelated-histories", false, "allow merging unrelated histories")
	pullCmd.Flags().BoolP("append", "a", false, "append to .git/FETCH_HEAD instead of overwriting")
	pullCmd.Flags().Bool("autostash", false, "automatically stash/stash pop before and after")
	pullCmd.Flags().String("cleanup", "", "how to strip spaces and #comments from message")
	pullCmd.Flags().Bool("commit", false, "perform a commit if the merge succeeds (default)")
	pullCmd.Flags().String("deepen", "", "deepen history of shallow clone")
	pullCmd.Flags().String("depth", "", "deepen history of shallow clone")
	pullCmd.Flags().Bool("dry-run", false, "dry run")
	pullCmd.Flags().Bool("edit", false, "edit message before committing")
	pullCmd.Flags().Bool("ff", false, "allow fast-forward")
	pullCmd.Flags().Bool("ff-only", false, "abort if fast-forward is not possible")
	pullCmd.Flags().BoolP("force", "f", false, "force overwrite of local branch")
	pullCmd.Flags().StringP("gpg-sign", "S", "", "GPG sign commit")
	pullCmd.Flags().BoolP("ipv4", "4", false, "use IPv4 addresses only")
	pullCmd.Flags().BoolP("ipv6", "6", false, "use IPv6 addresses only")
	pullCmd.Flags().StringP("jobs", "j", "", "number of submodules pulled in parallel")
	pullCmd.Flags().BoolP("keep", "k", false, "keep downloaded pack")
	pullCmd.Flags().String("log", "", "add (at most <n>) entries from shortlog to merge commit message")
	pullCmd.Flags().BoolS("n", "n", false, "do not show a diffstat at the end of the merge")
	pullCmd.Flags().String("negotiation-tip", "", "report that we have only objects reachable from this object")
	pullCmd.Flags().Bool("progress", false, "force progress reporting")
	pullCmd.Flags().BoolP("prune", "p", false, "prune remote-tracking branches no longer on remote")
	pullCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	pullCmd.Flags().StringP("rebase", "r", "", "incorporate changes by rebasing rather than merging")
	pullCmd.Flags().String("recurse-submodules", "", "control for recursive fetching of submodules")
	pullCmd.Flags().String("refmap", "", "specify fetch refmap")
	pullCmd.Flags().StringP("server-option", "o", "", "option to transmit")
	pullCmd.Flags().Bool("set-upstream", false, "set upstream for git pull/fetch")
	pullCmd.Flags().String("shallow-exclude", "", "deepen history of shallow clone, excluding rev")
	pullCmd.Flags().String("shallow-since", "", "deepen history of shallow repository based on time")
	pullCmd.Flags().Bool("show-forced-updates", false, "check for forced-updates on all updated branches")
	pullCmd.Flags().String("signoff", "", "add Signed-off-by:")
	pullCmd.Flags().Bool("squash", false, "create a single commit instead of doing a merge")
	pullCmd.Flags().Bool("stat", false, "show a diffstat at the end of the merge")
	pullCmd.Flags().StringP("strategy", "s", "", "merge strategy to use")
	pullCmd.Flags().StringP("strategy-option", "X", "", "option for selected merge strategy")
	pullCmd.Flags().BoolP("tags", "t", false, "fetch all tags and associated objects")
	pullCmd.Flags().Bool("unshallow", false, "convert to a complete repository")
	pullCmd.Flags().Bool("update-shallow", false, "accept refs that update .git/shallow")
	pullCmd.Flags().String("upload-pack", "", "path to upload pack on remote end")
	pullCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	pullCmd.Flags().Bool("verify-signatures", false, "verify that the named commit has a valid GPG signature")
	rootCmd.AddCommand(pullCmd)

	carapace.Gen(pullCmd).FlagCompletion(carapace.ActionMap{
		"cleanup": carapace.ActionValuesDescribed(
			"default", "act as 'strip' if the message is to be edited and as 'whitespace' otherwise",
			"scissors", "same as whitespace but cut from scissor line",
			"strip", "remove both whitespace and commentary lines",
			"verbatim", "don't change the commit message at all",
			"whitespace", "remove leading and trailing whitespace lines",
		),
		"rebase": carapace.ActionValuesDescribed(
			"false", "merge after fetching",
			"interactive", "allow list of commits to be edited",
			"merges", "try to rebase merges instead of skipping them",
			"preserve", "rebase and preserve merges",
			"true", "rebase after fetching",
		),
		"recurse-submodules": carapace.ActionValuesDescribed(
			"no", "disable recursion",
			"on-demand", "only when submodule reference in superproject is updated",
			"yes", "always recurse",
		).StyleF(style.ForKeyword),
		"strategy": carapace.ActionValues("octopus", "ours", "recursive", "resolve", "subtree"),
	})

	carapace.Gen(pullCmd).PositionalCompletion(
		git.ActionRemotes(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return git.ActionRemoteBranches(c.Args[0])
		}),
	)
}
