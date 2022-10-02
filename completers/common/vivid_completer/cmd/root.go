package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "vivid",
	Short: "LS_COLORS manager with multiple themes",
	Long:  "https://github.com/sharkdp/vivid",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringP("color-mode", "m", "", "Type of ANSI colors to be used")
	rootCmd.Flags().StringP("database", "d", "", "Path to filetypes database (filetypes.yml)")
	rootCmd.Flags().BoolP("help", "h", false, "Prints help information")
	rootCmd.Flags().BoolP("version", "V", false, "Prints version information")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"color-mode": carapace.ActionValues("8-bit", "24-bit"),
		"database":   carapace.ActionFiles(),
	})
}
