package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gitk",
	Short: "The Git repository browser",
	Long:  "https://git-scm.com/docs/gitk",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("all", false, "Show all refs (branches, tags, etc.).")
	rootCmd.Flags().Bool("ancestry-path", false, "only display commits that exist directly on the ancestry chain between the commit1 and commit2")
	rootCmd.Flags().String("argscmd", "", "Command to be run each time gitk has to determine the revision range to show")
	rootCmd.Flags().String("branches", "", "Pretend as if all the branches are listed on the command line")
	rootCmd.Flags().Bool("date-order", false, "Sort commits by date when possible.")
	rootCmd.Flags().Bool("full-history", false, "When filtering history with <path>..., does not prune some history")
	rootCmd.Flags().Bool("left-right", false, "Mark which side of a symmetric difference a commit is reachable from")
	rootCmd.Flags().Bool("merge", false, "After merge stops with conflicts, show the commits on the history between two branches")
	rootCmd.Flags().String("remotes", "", "Pretend as if all the remote branches are listed on the command line")
	rootCmd.Flags().String("select-commit", "", "Select the specified commit after loading the graph")
	rootCmd.Flags().Bool("simplify-merges", false, "Additional option to --full-history to remove some needless merges from the resulting history")
	rootCmd.Flags().String("since", "", "Show commits more recent than a specific date.")
	rootCmd.Flags().String("tags", "", "Pretend as if all the tags are listed on the command line")
	rootCmd.Flags().String("until", "", "Show commits older than a specific date.")

	rootCmd.Flag("argscmd").NoOptDefVal = " "
	rootCmd.Flag("branches").NoOptDefVal = " "
	rootCmd.Flag("remotes").NoOptDefVal = " "
	rootCmd.Flag("select-commit").NoOptDefVal = " "
	rootCmd.Flag("since").NoOptDefVal = " "
	rootCmd.Flag("tags").NoOptDefVal = " "
	rootCmd.Flag("until").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"argscmd":       carapace.ActionFiles(),
		"select-commit": git.ActionRefs(git.RefOption{}.Default()),
		"since":         time.ActionDate(),
		"until":         time.ActionDate(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionMultiParts("...", func(c carapace.Context) carapace.Action {
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)

	carapace.Gen(rootCmd).DashAnyCompletion(
		carapace.ActionFiles(),
	)
}
