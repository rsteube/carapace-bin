package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var generateRgbBreatheTableCmd = &cobra.Command{
	Use:   "generate-rgb-breathe-table",
	Short: "Generates an RGB Light breathing table header",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateRgbBreatheTableCmd).Standalone()

	generateRgbBreatheTableCmd.Flags().StringP("center", "c", "", "The breathing center value")
	generateRgbBreatheTableCmd.Flags().BoolP("help", "h", false, "show this help message and exit")
	generateRgbBreatheTableCmd.Flags().StringP("max", "m", "", "The breathing maximum value")
	generateRgbBreatheTableCmd.Flags().StringP("output", "o", "", "File to write to")
	generateRgbBreatheTableCmd.Flags().BoolP("quiet", "q", false, "Quiet mode, only output error messages")
	rootCmd.AddCommand(generateRgbBreatheTableCmd)

	carapace.Gen(generateRgbBreatheTableCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
