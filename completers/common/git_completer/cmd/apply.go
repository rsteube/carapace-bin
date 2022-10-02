package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply a patch to files and/or to the index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applyCmd).Standalone()

	applyCmd.Flags().BoolP("3way", "3", false, "attempt three-way merge, fall back on normal patch if that fails")
	applyCmd.Flags().StringS("C", "C", "", "ensure at least <n> lines of context match")
	applyCmd.Flags().Bool("allow-overlap", false, "allow overlapping hunks")
	applyCmd.Flags().Bool("apply", false, "also apply the patch (use with --stat/--summary/--check)")
	applyCmd.Flags().String("build-fake-ancestor", "", "build a temporary index based on embedded index information")
	applyCmd.Flags().Bool("cached", false, "apply a patch without touching the working tree")
	applyCmd.Flags().Bool("check", false, "instead of applying the patch, see if the patch is applicable")
	applyCmd.Flags().String("directory", "", "prepend <root> to all filenames")
	applyCmd.Flags().String("exclude", "", "don't apply changes matching the given path")
	applyCmd.Flags().Bool("ignore-space-change", false, "ignore changes in whitespace when finding context")
	applyCmd.Flags().Bool("ignore-whitespace", false, "ignore changes in whitespace when finding context")
	applyCmd.Flags().Bool("inaccurate-eof", false, "tolerate incorrectly detected missing new-line at the end of file")
	applyCmd.Flags().String("include", "", "apply changes matching the given path")
	applyCmd.Flags().Bool("index", false, "make sure the patch is applicable to the current index")
	applyCmd.Flags().BoolP("intent-to-add", "N", false, "mark new files with `git add --intent-to-add`")
	applyCmd.Flags().Bool("no-add", false, "ignore additions made by the patch")
	applyCmd.Flags().Bool("numstat", false, "show number of added and deleted lines in decimal notation")
	applyCmd.Flags().StringS("p", "p", "", "remove <num> leading slashes from traditional diff paths")
	applyCmd.Flags().Bool("recount", false, "do not trust the line counts in the hunk headers")
	applyCmd.Flags().Bool("reject", false, "leave the rejected hunks in corresponding *.rej files")
	applyCmd.Flags().BoolP("reverse", "R", false, "apply the patch in reverse")
	applyCmd.Flags().Bool("stat", false, "instead of applying the patch, output diffstat for the input")
	applyCmd.Flags().Bool("summary", false, "instead of applying the patch, output a summary for the input")
	applyCmd.Flags().Bool("unidiff-zero", false, "don't expect at least one line of context")
	applyCmd.Flags().Bool("unsafe-paths", false, "accept a patch that touches outside the working area")
	applyCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	applyCmd.Flags().String("whitespace", "", "detect new or modified lines that have whitespace errors")
	applyCmd.Flags().BoolS("z", "z", false, "paths are separated with NUL character")
	rootCmd.AddCommand(applyCmd)

	carapace.Gen(applyCmd).FlagCompletion(carapace.ActionMap{
		"build-fake-ancestor": carapace.ActionFiles(),
		"directory":           carapace.ActionDirectories(),
		"exclude":             carapace.ActionFiles(),
		"include":             carapace.ActionFiles(),
		"whitespace":          git.ActionWhitespaceModes(),
	})

	carapace.Gen(applyCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
