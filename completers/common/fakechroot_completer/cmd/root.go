package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fakechroot",
	Short: "gives a fake chroot environment",
	Long:  "https://linux.die.net/man/1/fakechroot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("bindir", "b", "", "Specify a directory which contains the replacement executables shipped with fakechroot.")
	rootCmd.Flags().StringP("config-dir", "c", "", "Specify a directory which contains additional configuration for fakechroot.")
	rootCmd.Flags().StringP("elfloader", "d", "", "Specify an alternative dynamic linker.")
	rootCmd.Flags().StringP("environment", "e", "", "Load additional configuration with environment.")
	rootCmd.Flags().BoolP("help", "h", false, "Display help.")
	rootCmd.Flags().StringP("lib", "l", "", "Specify an alternative wrapper library.")
	rootCmd.Flags().BoolP("use-system-libs", "s", false, "Use system libraries before chroot's libraries.")
	rootCmd.Flags().BoolP("version", "v", false, "Display version.")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bindir":     carapace.ActionDirectories(),
		"config-dir": carapace.ActionDirectories(),
	})
}
