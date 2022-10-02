package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bats",
	Short: "Bash Automated Testing System",
	Long:  "https://github.com/bats-core/bats-core",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("count", "c", false, "Count the number of test cases without running any tests")
	rootCmd.Flags().BoolP("help", "h", false, "Display this help message")
	rootCmd.Flags().BoolP("pretty", "p", false, "Show results in pretty format (default for terminals)")
	rootCmd.Flags().BoolP("tap", "t", false, "Show results in TAP format")
	rootCmd.Flags().BoolP("version", "v", false, "Display the version number")

	carapace.Gen(rootCmd).PositionalAnyCompletion(carapace.ActionFiles())
}
