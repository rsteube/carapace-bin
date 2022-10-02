package cmd

import (
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/ps"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "htop",
	Short: "interactive process viewer",
	Long:  "https://htop.dev/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("delay", "d", "", "Set the delay between updates, in tenths of seconds")
	rootCmd.Flags().String("drop-capabilities", "", "Drop Linux capabilities when running as root")
	rootCmd.Flags().StringP("filter", "F", "", "Show only the commands matching the given filter")
	rootCmd.Flags().BoolP("help", "h", false, "Print this help screen")
	rootCmd.Flags().BoolP("highlight-changes", "H", false, "Highlight new and old processes")
	rootCmd.Flags().BoolP("no-color", "C", false, "Use a monochrome color scheme")
	rootCmd.Flags().BoolP("no-mouse", "m", false, "Disable the mouse")
	rootCmd.Flags().BoolP("no-unicode", "U", false, "Do not use unicode but plain ASCII")
	rootCmd.Flags().StringP("pid", "p", "", "Show only the given PIDs")
	rootCmd.Flags().Bool("readonly", false, "Disable all system and process changing feature")
	rootCmd.Flags().StringP("sort-key", "s", "", "Sort by COLUMN (try --sort-key=help for a list)")
	rootCmd.Flags().BoolP("tree", "t", false, "Show the tree view by default")
	rootCmd.Flags().StringP("user", "u", "", "Show only processes for a given user (or $USER)")
	rootCmd.Flags().BoolP("version", "v", false, "Print version info")

	rootCmd.Flag("drop-capabilities").NoOptDefVal = " "
	rootCmd.Flag("highlight-changes").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"drop-capabilities": carapace.ActionValuesDescribed(
			"off", "do not drop any capabilities",
			"basic", "drop all capabilities not needed by htop",
			"strict", "drop all capabilities except those needed for core functionality",
		),
		"pid": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return ps.ActionProcessIds().Invoke(c).Filter(c.Parts).ToA()
		}),
		"sort-key": ActionSortKeys(),
		"user":     os.ActionUsers(),
	})
}

func ActionSortKeys() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("htop", "--sort-key", "help")(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				if fields := strings.Fields(line); len(fields) > 1 {
					vals = append(vals, fields[0], strings.Join(fields[1:], " "))
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
