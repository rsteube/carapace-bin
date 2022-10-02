package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "baobab",
	Short: "A graphical disk usage analyzer for the GNOME deskto",
	Long:  "https://wiki.gnome.org/action/show/Apps/DiskUsageAnalyzer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("all-file-systems", "a", false, "Do not skip directories on different file systems. Ignored if DIRECTORY is not specified.")
	rootCmd.Flags().BoolP("help", "h", false, "Show help options")
	rootCmd.Flags().BoolP("version", "v", false, "Print version information and exit")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
