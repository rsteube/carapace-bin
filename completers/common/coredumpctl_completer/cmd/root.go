package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/os"
	"github.com/rsteube/carapace-bin/pkg/actions/time"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "coredumpctl",
	Short: "List or retrieve coredumps from the journal",
	Long:  "https://www.freedesktop.org/software/systemd/man/coredumpctl.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("1", "1", false, "Show information about most recent entry only")
	rootCmd.Flags().Bool("all", false, "Look at all journal files instead of local ones")
	rootCmd.Flags().String("debugger", "", "Use the given debugger")
	rootCmd.Flags().StringP("debugger-arguments", "A", "", "Pass the given arguments to the debugger")
	rootCmd.Flags().StringP("directory", "D", "", "Use journal files from directory")
	rootCmd.Flags().StringP("field", "F", "", "List all values a certain field takes")
	rootCmd.Flags().String("file", "", "Use journal file")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().String("json", "", "Generate JSON output")
	rootCmd.Flags().StringS("n", "n", "", "Show maximum number of rows")
	rootCmd.Flags().Bool("no-legend", false, "Do not print the column headers")
	rootCmd.Flags().Bool("no-pager", false, "Do not pipe output into a pager")
	rootCmd.Flags().StringP("output", "o", "", "Write output to FILE")
	rootCmd.Flags().BoolP("quiet", "q", false, "Do not show info messages and privilege warning")
	rootCmd.Flags().BoolP("reverse", "r", false, "Show the newest entries first")
	rootCmd.Flags().StringP("since", "S", "", "Only print coredumps since the date")
	rootCmd.Flags().StringP("until", "U", "", "Only print coredumps until the date")
	rootCmd.Flags().Bool("version", false, "Print version string")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"debugger": carapace.Batch(
			os.ActionPathExecutables(),
			carapace.ActionFiles(),
		).ToA(),
		"directory": carapace.ActionDirectories(),
		"json":      carapace.ActionValues("pretty", "short", "off"),
		"output":    carapace.ActionFiles(),
		"since":     time.ActionDate(),
		"until":     time.ActionDate(),
	})
}
