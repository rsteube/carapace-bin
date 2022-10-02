package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze toit files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analyzeCmd).Standalone()
	analyzeCmd.Flags().String("sdk", "", "the default SDK path to use")
	analyzeCmd.Flags().String("toitc", "", "the default toit compiler to use")
	analyzeCmd.Flags().BoolP("verbose", "v", false, "")
	rootCmd.AddCommand(analyzeCmd)

	carapace.Gen(analyzeCmd).FlagCompletion(carapace.ActionMap{
		"sdk":   carapace.ActionDirectories(),
		"toitc": carapace.ActionFiles(),
	})

	carapace.Gen(analyzeCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
