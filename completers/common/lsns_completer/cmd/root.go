package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/lsns_completer/cmd/action"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsns",
	Short: "List system namespaces",
	Long:  "https://man7.org/linux/man-pages/man8/lsns.8.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("help", "h", false, "display this help")
	rootCmd.Flags().BoolP("json", "J", false, "use JSON output format")
	rootCmd.Flags().BoolP("list", "l", false, "use list format output")
	rootCmd.Flags().BoolP("noheadings", "n", false, "don't print headings")
	rootCmd.Flags().BoolP("notruncate", "u", false, "don't truncate text in columns")
	rootCmd.Flags().BoolP("nowrap", "W", false, "don't use multi-line representation")
	rootCmd.Flags().StringP("output", "o", "", "define which output columns to use")
	rootCmd.Flags().Bool("output-all", false, "output all columns")
	rootCmd.Flags().BoolP("raw", "r", false, "use the raw output format")
	rootCmd.Flags().StringP("task", "p", "", "print process namespaces")
	rootCmd.Flags().StringP("type", "t", "", "namespace type (mnt, net, ipc, user, pid, uts, cgroup, time)")
	rootCmd.Flags().BoolP("version", "V", false, "display version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionColumns().Invoke(c).Filter(c.Parts).ToA()
		}),
		"task": ps.ActionProcessIds(),
		"type": carapace.ActionValues("mnt", "net", "ipc", "user", "pid", "uts", "cgroup", "time"),
	})
}
