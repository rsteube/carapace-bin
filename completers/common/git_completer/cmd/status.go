package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "show the working tree status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(statusCmd).Standalone()

	statusCmd.Flags().Bool("ahead-behind", false, "display detailed ahead/behind counts for the branch relative to its upstream branch")
	statusCmd.Flags().BoolP("branch", "b", false, "Show the branch and tracking info even in short-format.")
	statusCmd.Flags().String("column", "", "Display untracked files in columns.")
	statusCmd.Flags().String("find-renames", "", "turn on rename detection, optionally setting the similarity threshold")
	statusCmd.Flags().String("ignore-submodules", "", "Ignore changes to submodules when looking for changes.")
	statusCmd.Flags().String("ignored", "", "Show ignored files as well")
	statusCmd.Flags().Bool("long", false, "Give the output in the long-format. This is the default.")
	statusCmd.Flags().Bool("no-ahead-behind", false, "do not display detailed ahead/behind counts for the branch relative to its upstream branch")
	statusCmd.Flags().Bool("no-column", false, "Do not display untracked files in columns.")
	statusCmd.Flags().Bool("no-renames", false, "turn off rename detection")
	statusCmd.Flags().String("porcelain", "", "Give the output in an easy-to-parse format for scripts.")
	statusCmd.Flags().Bool("renames", false, "turn on rename detection")
	statusCmd.Flags().BoolP("short", "s", false, "Give the output in the short-format.")
	statusCmd.Flags().Bool("show-stash", false, "Show the number of entries currently stashed away.")
	statusCmd.Flags().StringP("untracked-files", "u", "", "Show untracked files")
	statusCmd.Flags().BoolP("verbose", "v", false, "also show the textual changes")
	statusCmd.Flags().BoolS("z", "z", false, "Terminate entries with NUL, instead of LF.")
	rootCmd.AddCommand(statusCmd)

	statusCmd.Flag("column").NoOptDefVal = " "
	statusCmd.Flag("find-renames").NoOptDefVal = " "
	statusCmd.Flag("ignored").NoOptDefVal = "traditional"
	statusCmd.Flag("porcelain").NoOptDefVal = "v1"
	statusCmd.Flag("ignore-submodules").NoOptDefVal = "all"

	carapace.Gen(statusCmd).FlagCompletion(carapace.ActionMap{
		"column":            carapace.ActionValues("always", "never").StyleF(style.ForKeyword),
		"ignore-submodules": carapace.ActionValues("none", "untracked", "dirty", "all").StyleF(style.ForKeyword),
		"ignored":           ActionIgnoredModes(),
		"porcelain":         carapace.ActionValues("v1"),
		"untracked-files":   ActionUntrackedFilesModes(),
	})

	carapace.Gen(statusCmd).PositionalAnyCompletion(carapace.ActionFiles())
}

func ActionUntrackedFilesModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"no", "show no untracked files",
		"normal", "shows untracked files and directories",
		"all", "also shows individual files in untracked directories",
	)
}

func ActionIgnoredModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"traditional", "shows ignored files and directories",
		"no", "show no ignored files",
		"matching", "shows ignored files and directories matching an ignore pattern",
	)
}
