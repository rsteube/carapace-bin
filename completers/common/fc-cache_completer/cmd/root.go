package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fc-cache",
	Short: "Build font information caches",
	Long:  "https://linux.die.net/man/1/fc-cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("error-on-no-fonts", "E", false, "raise an error if no fonts in a directory")
	rootCmd.Flags().BoolP("force", "f", false, "scan directories with apparently valid caches")
	rootCmd.Flags().BoolP("help", "h", false, "display this help and exit")
	rootCmd.Flags().BoolP("really-force", "r", false, "erase all existing caches, then rescan")
	rootCmd.Flags().StringP("sysroot", "y", "", "prepend SYSROOT to all paths for scanning")
	rootCmd.Flags().BoolP("system-only", "s", false, "scan system-wide directories only")
	rootCmd.Flags().BoolP("verbose", "v", false, "display status information while busy")
	rootCmd.Flags().BoolP("version", "V", false, "display font config version and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
