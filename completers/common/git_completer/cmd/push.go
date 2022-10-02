package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Update remote refs along with associated objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pushCmd).Standalone()

	pushCmd.Flags().Bool("all", false, "push all refs")
	pushCmd.Flags().Bool("atomic", false, "request atomic transaction on remote side")
	pushCmd.Flags().BoolP("delete", "d", false, "delete refs")
	pushCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	pushCmd.Flags().String("exec", "", "receive pack program")
	pushCmd.Flags().Bool("follow-tags", false, "push missing but relevant tags")
	pushCmd.Flags().BoolP("force", "f", false, "force updates")
	pushCmd.Flags().String("force-with-lease", "", "require old value of ref to be at this value")
	pushCmd.Flags().BoolP("ipv4", "4", false, "use IPv4 addresses only")
	pushCmd.Flags().BoolP("ipv6", "6", false, "use IPv6 addresses only")
	pushCmd.Flags().Bool("mirror", false, "mirror all refs")
	pushCmd.Flags().Bool("no-verify", false, "bypass pre-push hook")
	pushCmd.Flags().Bool("porcelain", false, "machine-readable output")
	pushCmd.Flags().Bool("progress", false, "force progress reporting")
	pushCmd.Flags().Bool("prune", false, "prune locally removed refs")
	pushCmd.Flags().StringP("push-option", "o", "", "option to transmit")
	pushCmd.Flags().BoolP("quiet", "q", false, "be more quiet")
	pushCmd.Flags().String("receive-pack", "", "receive pack program")
	pushCmd.Flags().String("recurse-submodules", "", "control recursive pushing of submodules")
	pushCmd.Flags().String("repo", "", "repository")
	pushCmd.Flags().BoolP("set-upstream", "u", false, "set upstream for git pull/status")
	pushCmd.Flags().String("signed", "", "GPG sign the push")
	pushCmd.Flags().Bool("tags", false, "push tags (can't be used with --all or --mirror)")
	pushCmd.Flags().Bool("thin", false, "use thin pack")
	pushCmd.Flags().BoolP("verbose", "v", false, "be more verbose")
	rootCmd.AddCommand(pushCmd)

	pushCmd.Flag("force-with-lease").NoOptDefVal = " "

	carapace.Gen(pushCmd).FlagCompletion(carapace.ActionMap{
		"signed": carapace.ActionValues("yes", "no", "if-asked"),
	})

	carapace.Gen(pushCmd).PositionalCompletion(
		git.ActionRemotes(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if pushCmd.Flag("set-upstream").Changed {
				// if set-upstream is set the desired remote branch is likely named the same as the current
				return git.ActionCurrentBranch()
			} else {
				return git.ActionRemoteBranches(c.Args[0])
			}
		}),
	)
}
