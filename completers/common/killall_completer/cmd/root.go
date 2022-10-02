package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "killall",
	Short: "kill processes by name",
	Long:  "https://linux.die.net/man/1/killall",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("exact", "e", false, "require exact match for very long names")
	rootCmd.Flags().BoolP("ignore-case", "I", false, "case insensitive process name match")
	rootCmd.Flags().BoolP("interactive", "i", false, "ask for confirmation before killing")
	rootCmd.Flags().BoolP("list", "l", false, "list all known signal names")
	rootCmd.Flags().StringP("ns", "n", "", "match processes that belong to the same namespaces as PID")
	rootCmd.Flags().StringP("older-than", "o", "", "kill processes older than TIME")
	rootCmd.Flags().BoolP("process-group", "g", false, "kill process group instead of process")
	rootCmd.Flags().BoolP("quiet", "q", false, "don't print complaints")
	rootCmd.Flags().BoolP("regexp", "r", false, "interpret NAME as an extended regular expression")
	rootCmd.Flags().StringP("signal", "s", "", "send this signal instead of SIGTERM")
	rootCmd.Flags().StringP("user", "u", "", "kill only process(es) running as USER")
	rootCmd.Flags().BoolP("verbose", "v", false, "report if the signal was successfully sent")
	rootCmd.Flags().BoolP("version", "V", false, "display version information")
	rootCmd.Flags().BoolP("wait", "w", false, "wait for processes to die")
	rootCmd.Flags().StringP("younger-than", "y", "", "kill processes younger than TIME")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"ns":     ps.ActionProcessIds(),
		"signal": ps.ActionKillSignals(),
		"user":   os.ActionUsers(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		os.ActionPathExecutables(),
	)
}
