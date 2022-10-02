package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/git"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/npm"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/turbo"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run tasks across projects in your monorepo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().String("cache-dir", "", "specify local filesystem cache directory")
	runCmd.Flags().String("concurrency", "", "set/limit the max concurrency of task execution")
	runCmd.Flags().Bool("continue", false, "whether or not to continue with execution in the presence of an error")
	runCmd.Flags().String("cpuprofile", "", "outputs the cpuprofile to the given file")
	runCmd.Flags().String("cwd", "", "set the working directory of the command")
	runCmd.Flags().String("dry", "", "instead of executing tasks, display details about the affected packages")
	runCmd.Flags().String("dry-run", "", "instead of executing tasks, display details about the affected packages")
	runCmd.Flags().StringArray("filter", []string{}, "specify combinations of package(s)/apps, directories, and git commits to act as entrypoints for execution")
	runCmd.Flags().Bool("force", false, "ignore existing cached artifacts and forcibly re-execute all tasks")
	runCmd.Flags().StringArray("global-deps", []string{}, "specify glob of global filesystem dependencies to be hashed")
	runCmd.Flags().String("graph", "", "generate graph of the current task")
	runCmd.Flags().String("heap", "", "outputs the trace to the given file")
	runCmd.Flags().StringArray("ignore", []string{}, "ignore files or directories from impacting scope")
	runCmd.Flags().Bool("no-cache", false, "do not cache results of the task")
	runCmd.Flags().Bool("only", false, "restricts execution to only include specified tasks")
	runCmd.Flags().String("output-logs", "", "set type of process output logging")
	runCmd.Flags().Bool("parallel", false, "run commands in parallel across packages and apps and ignore the dependency graph")
	runCmd.Flags().Bool("preflight", false, "enables sending a preflight request before every cache artifact and analytics request")
	runCmd.Flags().Bool("remote-only", false, "ignore the local filesystem cache for all tasks")
	runCmd.Flags().String("team", "", "the slug of the remote cache team")
	runCmd.Flags().String("token", "", "a bearer token for remote caching")
	runCmd.Flags().String("trace", "", "outputs the trace to the given file")
	runCmd.Flags().CountP("v", "v", "specify log level")

	runCmd.Flag("cache-dir").NoOptDefVal = " "
	runCmd.Flag("concurrency").NoOptDefVal = " "
	runCmd.Flag("cwd").NoOptDefVal = " "
	runCmd.Flag("dry").NoOptDefVal = " "
	runCmd.Flag("dry-run").NoOptDefVal = " "
	runCmd.Flag("filter").NoOptDefVal = " "
	runCmd.Flag("graph").NoOptDefVal = " "
	runCmd.Flag("global-deps").NoOptDefVal = " "
	runCmd.Flag("ignore").NoOptDefVal = " "
	runCmd.Flag("output-logs").NoOptDefVal = " "
	runCmd.Flag("token").NoOptDefVal = " "
	runCmd.Flag("team").NoOptDefVal = " "
	runCmd.Flag("trace").NoOptDefVal = " "
	runCmd.Flag("heap").NoOptDefVal = " "
	runCmd.Flag("cpuprofile").NoOptDefVal = " "

	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"cache-dir":  carapace.ActionDirectories(),
		"cpuprofile": carapace.ActionFiles(),
		"cwd":        carapace.ActionDirectories(),
		"dry":        carapace.ActionValues("json"),
		"dry-run":    carapace.ActionValues("json"),
		"filter": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if index := strings.LastIndexAny(c.CallbackValue, "[]"); index != -1 && []rune(c.CallbackValue)[index] == '[' {
				return git.ActionRefs(git.RefOption{}.Default()).Invoke(c).Prefix(c.CallbackValue[:index+1]).Suffix("]").ToA().NoSpace()
			}
			if index := strings.LastIndexAny(c.CallbackValue, "{}"); index != -1 && []rune(c.CallbackValue)[index] == '{' {
				prefix := c.CallbackValue[:index+1]
				c.CallbackValue = c.CallbackValue[index+1:]
				return carapace.ActionDirectories().Invoke(c).Prefix(prefix).ToA().NoSpace()
			}
			return npm.ActionDependencies().NoSpace()
		}),
		"graph": carapace.ActionFiles(),
		"heap":  carapace.ActionFiles(),
		"output-logs": carapace.ActionValuesDescribed(
			"full", "show all output",
			"hash-only", "show only turbo-computed task hashes",
			"new-only", "show only new output with only hashes for cached tasks",
			"none", "hide process output",
		),
		"trace": carapace.ActionFiles(),
	})

	carapace.Gen(runCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return turbo.ActionPipelineTasks().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
