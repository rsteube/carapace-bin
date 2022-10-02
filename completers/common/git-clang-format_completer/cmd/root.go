package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "git-clang-format",
	Short: "run clang-format on lines that differ",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("Cached,", false, "format lines in the stage instead of the working dir")
	rootCmd.Flags().String("binary", "", "path to clang-format")
	rootCmd.Flags().String("commit", "", "default commit to use if none is specified")
	rootCmd.Flags().Bool("diff", false, "print a diff instead of applying the changes")
	rootCmd.Flags().Bool("diffstat", false, "print a diffstat instead of applying the changes")
	rootCmd.Flags().String("extensions", "", "comma-separated list of file extensions to format")
	rootCmd.Flags().BoolP("force", "f", false, "allow changes to unstaged files")
	rootCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	rootCmd.Flags().BoolP("patch", "p", false, "select hunks interactively")
	rootCmd.Flags().BoolP("quiet", "q", false, "print less information")
	rootCmd.Flags().Bool("staged,", false, "format lines in the stage instead of the working dir")
	rootCmd.Flags().String("style", "", "passed to clang-format")
	rootCmd.Flags().BoolP("verbose", "v", false, "print extra information")

	// TODO flag completion
	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"binary": carapace.ActionFiles(),
		"commit": git.ActionRefs(git.RefOption{}.Default()),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if rootCmd.Flags().ArgsLenAtDash() != -1 {
				return carapace.ActionFiles()
			}

			if diff := rootCmd.Flag("diff").Changed; len(c.Args) == 0 || (len(c.Args) == 1 && diff) {
				return git.ActionRefs(git.RefOption{}.Default())
			}
			return carapace.ActionFiles()
		}),
	)
}
