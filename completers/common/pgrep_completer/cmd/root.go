package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pgrep",
	Short: "look up processes based on name and other attributes",
	Long:  "https://linux.die.net/man/1/pgrep",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("count", "c", false, "count of matching processes")
	rootCmd.Flags().StringP("delimiter", "d", "", "specify output delimiter")
	rootCmd.Flags().StringP("euid", "u", "", "match by effective IDs")
	rootCmd.Flags().BoolP("exact", "x", false, "match exactly with the command name")
	rootCmd.Flags().BoolP("full", "f", false, "use full process name to match")
	rootCmd.Flags().StringP("group", "G", "", "match real group IDs")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("ignore-case", "i", false, "match case insensitively")
	rootCmd.Flags().BoolP("inverse", "v", false, "negates the matching")
	rootCmd.Flags().BoolP("lightweight", "w", false, "list all TID")
	rootCmd.Flags().BoolP("list-full", "a", false, "list PID and full command line")
	rootCmd.Flags().BoolP("list-name", "l", false, "list PID and process name")
	rootCmd.Flags().BoolP("logpidfile", "L", false, "fail if PID file is not locked")
	rootCmd.Flags().BoolP("newest", "n", false, "select most recently started")
	rootCmd.Flags().String("ns", "", "match the processes that belong to the same namespace as <pid>")
	rootCmd.Flags().String("nslist", "", "list which namespaces will be considered for the --ns option.")
	rootCmd.Flags().StringP("older", "O", "", "select where older than seconds")
	rootCmd.Flags().BoolP("oldest", "o", false, "select least recently started")
	rootCmd.Flags().StringP("parent", "P", "", "match only child processes of the given parent")
	rootCmd.Flags().StringP("pgroup", "g", "", "match listed process group IDs")
	rootCmd.Flags().StringP("pidfile", "F", "", "read PIDs from file")
	rootCmd.Flags().StringP("runstates", "r", "", "match runstates")
	rootCmd.Flags().StringP("session", "s", "", "match session IDs")
	rootCmd.Flags().StringP("terminal", "t", "", "match by controlling terminal")
	rootCmd.Flags().StringP("uid", "U", "", "match by real IDs")
	rootCmd.Flags().BoolP("version", "V", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"euid": os.ActionUsers(),
		"group": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionGroups().Invoke(c).Filter(c.Parts).ToA()
		}),
		"ns": ps.ActionProcessIds(),
		"nslist": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ipc", "mnt", "net", "pid", "user", "uts").Invoke(c).Filter(c.Parts).ToA()
		}),
		"parent": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ps.ActionProcessIds().Invoke(c).Filter(c.Parts).ToA()
		}),
		"pidfile": carapace.ActionFiles(),
		"runstates": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ps.ActionProcessStates().Invoke(c).Filter(c.Parts).ToA()
		}),
		"session": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionSessionIds().Invoke(c).Filter(c.Parts).ToA()
		}),
		"terminal": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return os.ActionTerminals().Invoke(c).Filter(c.Parts).ToA()
		}),
		"uid": os.ActionUsers(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		ps.ActionProcessExecutables(),
	)
}
