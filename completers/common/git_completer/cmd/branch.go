package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "List, create, or delete branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(branchCmd).Standalone()

	branchCmd.Flags().BoolS("C", "C", false, "Shortcut for --copy --force.")
	branchCmd.Flags().BoolS("D", "D", false, "Shortcut for --delete --force.")
	branchCmd.Flags().BoolS("M", "M", false, "Shortcut for --move --force.")
	branchCmd.Flags().String("abbrev", "", "Alter the sha1s minimum display length in the output listing.")
	branchCmd.Flags().BoolP("all", "a", false, "List both remote-tracking branches and local branches.")
	branchCmd.Flags().String("color", "", "Color branches to highlight current, local, and remote-tracking branches.")
	branchCmd.Flags().String("contains", "", "Only list branches which contain the specified commit (HEAD if not specified).")
	branchCmd.Flags().BoolP("copy", "c", false, "Copy a branch and the corresponding reflog.")
	branchCmd.Flags().Bool("create-reflog", false, "Create the branchs reflog.")
	branchCmd.Flags().BoolP("delete", "d", false, "Delete a branch.")
	branchCmd.Flags().Bool("edit-description", false, "Open an editor and edit the text to explain what the branch is for.")
	branchCmd.Flags().BoolP("force", "f", false, "Reset <branchname> to <startpoint>, even if <branchname> exists already.")
	branchCmd.Flags().String("format", "", "A string that interpolates %(fieldname) from a branch ref being shown and the object it points at.")
	branchCmd.Flags().BoolP("ignore-case", "i", false, "Sorting and filtering branches are case insensitive.")
	branchCmd.Flags().BoolP("list", "l", false, "List branches.")
	branchCmd.Flags().String("merged", "", "Only list branches whose tips are reachable from the specified commit (HEAD if not specified).")
	branchCmd.Flags().BoolP("move", "m", false, "Move/rename a branch and the corresponding reflog.")
	branchCmd.Flags().Bool("no-abbrev", false, "Display the full sha1s in the output listing rather than abbreviating them.")
	branchCmd.Flags().Bool("no-color", false, "Turn off branch colors, even when the configuration file gives the default to color output.")
	branchCmd.Flags().String("no-contains", "", "Only list branches which dont contain the specified commit (HEAD if not specified).")
	branchCmd.Flags().String("no-merged", "", "Only list branches whose tips are not reachable from the specified commit (HEAD if not specified).")
	branchCmd.Flags().Bool("no-track", false, "Do not set up upstream configuration, even if the branch.autoSetupMerge configuration variable is true.")
	branchCmd.Flags().String("points-at", "", "Only list branches of the given object.")
	branchCmd.Flags().BoolP("quiet", "q", false, "Be more quiet when creating or deleting a branch, suppressing non-error messages.")
	branchCmd.Flags().BoolP("remotes", "r", false, "List or delete (if used with -d) the remote-tracking branches.")
	branchCmd.Flags().String("set-upstream-to", "", "Set up <branchname>s tracking information so <upstream> is considered <branchname>s upstream branch.")
	branchCmd.Flags().Bool("show-current", false, "Print the name of the current branch.")
	branchCmd.Flags().String("sort", "", "Sort based on the key given.")
	branchCmd.Flags().BoolP("track", "t", false, "When creating a new branch, set up branch.<name>.remote and branch.<name>.merge configuration entries to mark the start-point branch as upstream from the new branch.")
	branchCmd.Flags().Bool("unset-upstream", false, "Remove the upstream information for <branchname>.")
	rootCmd.AddCommand(branchCmd)

	carapace.Gen(branchCmd).FlagCompletion(carapace.ActionMap{
		"D":               git.ActionRefs(git.RefOption{}.Default()),
		"color":           carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"contains":        git.ActionRefs(git.RefOption{}.Default()),
		"delete":          git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true}),
		"merged":          git.ActionRefs(git.RefOption{}.Default()),
		"no-contains":     git.ActionRefs(git.RefOption{}.Default()),
		"no-merged":       git.ActionRefs(git.RefOption{}.Default()),
		"points-at":       git.ActionRefs(git.RefOption{RemoteBranches: true, Tags: true}),
		"set-upstream-to": git.ActionRefs(git.RefOption{RemoteBranches: true, Tags: true}),
		"sort":            git.ActionFieldNames(),
	})

	carapace.Gen(branchCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch len(c.Args) {
			case 0:
				if branchCmd.Flag("set-upstream-to").Changed ||
					branchCmd.Flag("track").Changed ||
					branchCmd.Flag("no-track").Changed ||
					branchCmd.Flag("unset-upstream").Changed ||
					branchCmd.Flag("M").Changed ||
					branchCmd.Flag("C").Changed ||
					branchCmd.Flag("D").Changed ||
					branchCmd.Flag("move").Changed ||
					branchCmd.Flag("copy").Changed ||
					branchCmd.Flag("delete").Changed ||
					branchCmd.Flag("edit-description").Changed {
					return git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true, Tags: true}).Invoke(c).Filter(c.Args).ToA()
				}
			case 1:
				if branchCmd.Flag("M").Changed ||
					branchCmd.Flag("C").Changed ||
					branchCmd.Flag("D").Changed ||
					branchCmd.Flag("move").Changed ||
					branchCmd.Flag("copy").Changed ||
					branchCmd.Flag("delete").Changed {
					return git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true, Tags: true}).Invoke(c).Filter(c.Args).ToA()
				}
			default:
				if branchCmd.Flag("D").Changed ||
					branchCmd.Flag("delete").Changed {
					return git.ActionRefs(git.RefOption{LocalBranches: true, RemoteBranches: true, Tags: true}).Invoke(c).Filter(c.Args).ToA()
				}
			}
			return carapace.ActionValues()
		}),
	)
}
