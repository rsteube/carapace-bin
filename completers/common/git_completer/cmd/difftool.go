package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

var difftoolCmd = &cobra.Command{
	Use:   "difftool",
	Short: "Show changes using common diff tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(difftoolCmd).Standalone()

	difftoolCmd.Flags().BoolP("dir-diff", "d", false, "perform a full-directory diff")
	difftoolCmd.Flags().StringP("extcmd", "x", "", "specify a custom command for viewing diffs")
	difftoolCmd.Flags().BoolP("gui", "g", false, "use `diff.guitool` instead of `diff.tool`")
	difftoolCmd.Flags().Bool("no-index", false, "passed to `diff`")
	difftoolCmd.Flags().BoolP("no-prompt", "y", false, "do not prompt before launching a diff tool")
	difftoolCmd.Flags().Bool("symlinks", false, "use symlinks in dir-diff mode")
	difftoolCmd.Flags().StringP("tool", "t", "", "use the specified diff tool")
	difftoolCmd.Flags().Bool("tool-help", false, "print a list of diff tools that may be used with `--tool`")
	difftoolCmd.Flags().Bool("trust-exit-code", false, "exit when an invoked diff tool returns a non - zero exit code")
	rootCmd.AddCommand(difftoolCmd)

	carapace.Gen(difftoolCmd).FlagCompletion(carapace.ActionMap{
		"tool": carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
	})

	carapace.Gen(difftoolCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if util.HasPathPrefix(c.CallbackValue) {
				return carapace.ActionFiles()
			}
			return git.ActionRefs(git.RefOption{}.Default())
		}),
	)

	carapace.Gen(difftoolCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)

	carapace.Gen(difftoolCmd).DashAnyCompletion(
		carapace.ActionFiles(),
	)
}
