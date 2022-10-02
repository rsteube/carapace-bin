package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "acpi",
	Short: "Shows information from the /proc filesystem",
	Long:  "https://www.commandlinux.com/man-page/man1/acpi.1.html",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("ac-adapter", "a", false, "ac adapter information")
	rootCmd.Flags().BoolP("battery", "b", false, "battery information")
	rootCmd.Flags().BoolP("cooling", "c", false, "cooling information")
	rootCmd.Flags().BoolP("details", "i", false, "show additional details if available:")
	rootCmd.Flags().StringP("directory", "d", "", "path to ACPI info (/sys/class resp. /proc/acpi)")
	rootCmd.Flags().BoolP("everything", "V", false, "show every device, overrides above options")
	rootCmd.Flags().BoolP("fahrenheit", "f", false, "use fahrenheit as the temperature unit")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("kelvin", "k", false, "use kelvin as the temperature unit")
	rootCmd.Flags().BoolP("proc", "p", false, "use old proc interface instead of new sys interface")
	rootCmd.Flags().BoolP("show-empty", "s", false, "show non-operational devices")
	rootCmd.Flags().BoolP("thermal", "t", false, "thermal information")
	rootCmd.Flags().BoolP("version", "v", false, "output version information and exit")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"directory": carapace.ActionDirectories(),
	})
}
