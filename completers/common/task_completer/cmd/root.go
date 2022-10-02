package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/task"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "A task runner / simpler Make alternative written in Go",
	Long:  "https://taskfile.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("color", "c", false, "colored output. Enabled by default. Set flag to false or use NO_COLOR=1 to disable (default true)")
	rootCmd.Flags().StringP("concurrency", "C", "", "limit number tasks to run concurrently")
	rootCmd.Flags().StringP("dir", "d", "", "sets directory of execution")
	rootCmd.Flags().BoolP("dry", "n", false, "compiles and prints tasks in the order that they would be run, without executing them")
	rootCmd.Flags().BoolP("exit-code", "x", false, "pass-through the exit code of the task command")
	rootCmd.Flags().BoolP("force", "f", false, "forces execution even when the task is up-to-date")
	rootCmd.Flags().BoolP("help", "h", false, "shows Task usage")
	rootCmd.Flags().BoolP("init", "i", false, "creates a new Taskfile.yaml in the current folder")
	rootCmd.Flags().BoolP("list", "l", false, "lists tasks with description of current Taskfile")
	rootCmd.Flags().BoolP("list-all", "a", false, "lists tasks with or without a description")
	rootCmd.Flags().StringP("output", "o", "", "sets output style: [interleaved|group|prefixed]")
	rootCmd.Flags().String("output-group-begin", "", "message template to print before a task's grouped output")
	rootCmd.Flags().String("output-group-end", "", "message template to print after a task's grouped output")
	rootCmd.Flags().BoolP("parallel", "p", false, "executes tasks provided on command line in parallel")
	rootCmd.Flags().BoolP("silent", "s", false, "disables echoing")
	rootCmd.Flags().Bool("status", false, "exits with non-zero exit code if any of the given tasks is not up-to-date")
	rootCmd.Flags().Bool("summary", false, "show summary about a task")
	rootCmd.Flags().StringP("taskfile", "t", "", "choose which Taskfile to run. Defaults to \"Taskfile.yml\"")
	rootCmd.Flags().BoolP("verbose", "v", false, "enables verbose mode")
	rootCmd.Flags().Bool("version", false, "show Task version")
	rootCmd.Flags().BoolP("watch", "w", false, "enables watch of the given task")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"dir":      carapace.ActionDirectories(),
		"output":   carapace.ActionValues("interleaved", "group", "prefixed"),
		"taskfile": carapace.ActionFiles(".yml", ".yaml"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return task.ActionTasks(rootCmd.Flag("taskfile").Value.String()).Chdir(rootCmd.Flag("dir").Value.String()).Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
