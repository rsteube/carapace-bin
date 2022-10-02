package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var blameCmd = &cobra.Command{
	Use:     "blame",
	Short:   "Show what revision and author last modified each line of a file",
	Aliases: []string{"annotate"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blameCmd).Standalone()

	blameCmd.Flags().StringS("C", "C", "", "find line copies within and across files")
	blameCmd.Flags().StringS("L", "L", "", "process only line range <start>,<end> or function :<funcname>")
	blameCmd.Flags().StringS("M", "M", "", "find line movements within and across files")
	blameCmd.Flags().StringS("S", "S", "", "use revisions from <file> instead of calling git-rev-list")
	blameCmd.Flags().String("abbrev", "", "use <n> digits to display object names")
	blameCmd.Flags().BoolS("b", "b", false, "do not show object names of boundary commits")
	blameCmd.Flags().BoolS("c", "c", false, "use the same output mode as git-annotate ")
	blameCmd.Flags().Bool("color-by-age", false, "color lines by age")
	blameCmd.Flags().Bool("color-lines", false, "color redundant metadata from previous line differently")
	blameCmd.Flags().String("contents", "", "use <file>'s contents as the final image")
	blameCmd.Flags().String("ignore-rev", "", "ignore <rev> when blaming")
	blameCmd.Flags().String("ignore-revs-file", "", "ignore revisions from <file>")
	blameCmd.Flags().Bool("incremental", false, "show blame entries as we find them, incrementally")
	blameCmd.Flags().BoolS("l", "l", false, "show long commit SHA1 ")
	blameCmd.Flags().Bool("line-porcelain", false, "show porcelain format with per-line commit information")
	blameCmd.Flags().Bool("minimal", false, "spend extra cycles to find better match")
	blameCmd.Flags().BoolP("porcelain", "p", false, "show in a format designed for machine consumption")
	blameCmd.Flags().Bool("progress", false, "force progress reporting")
	blameCmd.Flags().Bool("root", false, "do not treat root commits as boundaries")
	blameCmd.Flags().BoolS("s", "s", false, "suppress author name and timestamp ")
	blameCmd.Flags().Bool("score-debug", false, "show output score for blame entries")
	blameCmd.Flags().BoolP("show-email", "e", false, "show author email instead of name ")
	blameCmd.Flags().BoolP("show-name", "f", false, "show original filename ")
	blameCmd.Flags().BoolP("show-number", "n", false, "show original linenumber ")
	blameCmd.Flags().Bool("show-stats", false, "show work cost statistics")
	blameCmd.Flags().BoolS("t", "t", false, "show raw timestamp ")
	blameCmd.Flags().BoolS("w", "w", false, "ignore whitespace differences")

	blameCmd.Flag("abbrev").NoOptDefVal = " "
	blameCmd.Flag("C").NoOptDefVal = " "
	blameCmd.Flag("M").NoOptDefVal = " "

	rootCmd.AddCommand(blameCmd)

	carapace.Gen(blameCmd).FlagCompletion(carapace.ActionMap{
		"S":                carapace.ActionFiles(),
		"contents":         carapace.ActionFiles(),
		"ignore-rev":       git.ActionRefs(git.RefOption{}.Default()),
		"ignore-revs-file": carapace.ActionFiles(),
	})

	carapace.Gen(blameCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.Args[0]) {
				return carapace.ActionValues()
			}
			return carapace.ActionFiles()
		}),
	)
}
