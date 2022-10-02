package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yj",
	Short: "Convert between YAML, TOML, JSON, and HCL",
	Long:  "https://github.com/sclevine/yj",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("e", "e", false, "Escape HTML (JSON out only)")
	rootCmd.Flags().BoolS("h", "h", false, "Show this help message")
	rootCmd.Flags().BoolS("i", "i", false, "Indent output (JSON or TOML out only)")
	rootCmd.Flags().BoolS("k", "k", false, "Attempt to parse keys as objects or numbers types (YAML out only)")
	rootCmd.Flags().BoolS("n", "n", false, "Do not covert inf, -inf, and NaN to/from strings (YAML or TOML only)")
	rootCmd.Flags().BoolS("v", "v", false, "Show version")

	carapace.Gen(rootCmd).PositionalCompletion(
		ActionFormats(),
		ActionFormats(),
	)
}

func ActionFormats() carapace.Action {
	return carapace.ActionValuesDescribed(
		"c", "HCL",
		"j", "JSON",
		"t", "TOML",
		"y", "YAML",
	)
}
