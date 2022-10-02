package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var symbolizeCmd = &cobra.Command{
	Use:   "symbolize",
	Short: "Symbolize a stack trace from an AOT-compiled Flutter app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(symbolizeCmd).Standalone()

	symbolizeCmd.Flags().StringP("debug-info", "d", "", "A path to the symbols file generated with \"--split-debug-info\".")
	symbolizeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	symbolizeCmd.Flags().StringP("input", "i", "", "A file path containing a Dart stack trace.")
	symbolizeCmd.Flags().StringP("output", "o", "", "A file path for a symbolized stack trace to be written to.")
	rootCmd.AddCommand(symbolizeCmd)

	carapace.Gen(symbolizeCmd).FlagCompletion(carapace.ActionMap{
		"debug-info": carapace.ActionFiles(),
		"input":      carapace.ActionFiles(),
		"output":     carapace.ActionFiles(),
	})
}
