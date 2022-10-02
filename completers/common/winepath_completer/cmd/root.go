package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "winepath",
	Short: "Tool to convert Unix paths to/from Win32 paths",
	Long:  "https://wiki.winehq.org/Winepath",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("0", "0", false, "separate output with \\0 character, instead of a newline")
	rootCmd.Flags().BoolP("help", "h", false, "output this help message and exit")
	rootCmd.Flags().BoolP("long", "l", false, "converts the short Windows path of an existing file or directory to the long format")
	rootCmd.Flags().BoolP("short", "s", false, "converts the long Windows path of an existing file or directory to the short format")
	rootCmd.Flags().BoolP("unix", "u", false, "converts a Windows path to a Unix path")
	rootCmd.Flags().BoolP("windows", "w", false, "converts a Unix path to a long Windows path")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
