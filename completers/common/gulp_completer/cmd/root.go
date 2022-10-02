package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gulp_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gulp",
	Short: "Command Line Interface for gulp",
	Long:  "https://github.com/gulpjs/gulp-cli",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().CountS("L", "L", "Set the loglevel.")
	rootCmd.Flags().Bool("color", false, "Will force gulp and gulp plugins to display colors.")
	rootCmd.Flags().Bool("compact-tasks", false, "Reduce the output of task dependency tree.")
	rootCmd.Flags().Bool("continue", false, "Continue execution of tasks upon failure.")
	rootCmd.Flags().String("cwd", "", "Manually set the CWD.")
	rootCmd.Flags().String("depth", "", "Specify the depth of the task dependency tree.")
	rootCmd.Flags().StringP("gulpfile", "f", "", "Manually set path of gulpfile.")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help.")
	rootCmd.Flags().String("log-level", "", "Set the loglevel.")
	rootCmd.Flags().Bool("no-color", false, "Will force gulp and gulp plugins to not display colors.")
	rootCmd.Flags().String("require", "", "Will require a module before running the gulpfile.")
	rootCmd.Flags().Bool("series", false, "Run tasks given on the CLI in series.")
	rootCmd.Flags().BoolP("silent", "S", false, "Suppress all gulp logging.")
	rootCmd.Flags().Bool("sort-tasks", false, "Will sort top tasks of task dependency tree. ")
	rootCmd.Flags().BoolP("tasks", "T", false, "Print the task dependency tree for the loaded gulpfile.")
	rootCmd.Flags().String("tasks-depth", "", "Specify the depth of the task dependency tree.")
	rootCmd.Flags().Bool("tasks-json", false, "Print the task dependency tree, in JSON format.")
	rootCmd.Flags().Bool("tasks-simple", false, "Print a plaintext list of tasks for the loaded gulpfile.")
	rootCmd.Flags().Bool("verify", false, "Will verify plugins referenced in project's package.json.")
	rootCmd.Flags().BoolP("version", "v", false, "Print the global and local gulp versions.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cwd":      carapace.ActionDirectories(),
		"gulpfile": carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		action.ActionTasks(rootCmd),
	)
}
