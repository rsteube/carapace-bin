package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "rmdir",
	Short: "remove empty directories",
	Long:  "https://linux.die.net/man/1/rmdir",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("help", false, "display this help and exit")
	rootCmd.Flags().Bool("ignore-fail-on-non-empty", false, "each failure that is solely because a directory is non-empty")
	rootCmd.Flags().BoolP("parents", "p", false, "remove DIRECTORY and its ancestors")
	rootCmd.Flags().BoolP("verbose", "v", false, "output a diagnostic for every directory processed")
	rootCmd.Flags().Bool("version", false, "output version information and exit")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
