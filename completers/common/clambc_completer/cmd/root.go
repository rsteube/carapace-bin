package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "clambc",
	Short: "Bytecode Analysis and Testing Tool",
	Long:  "http://www.clamav.net/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("debug", false, "Show debug")
	rootCmd.Flags().BoolP("force-interpreter", "f", false, "Force using the interpreter instead of the JIT")
	rootCmd.Flags().BoolP("help", "h", false, "Show this help")
	rootCmd.Flags().BoolP("info", "i", false, "Print information about bytecode")
	rootCmd.Flags().BoolP("input", "r", false, "Input file to run the bytecode on")
	rootCmd.Flags().BoolP("no-trace-showsource", "s", false, "Don't show source line during tracing")
	rootCmd.Flags().BoolP("printbcir", "c", false, "Print IR of bytecode signature")
	rootCmd.Flags().BoolP("printsrc", "p", false, "Print bytecode source")
	rootCmd.Flags().String("statistics", "", "Collect and print bytecode execution statistics")
	rootCmd.Flags().StringP("trace", "T", "", "Set bytecode trace level 0..7 (default 7)")
	rootCmd.Flags().BoolP("trust-bytecode", "t", false, "Trust loaded bytecode (default yes)")
	rootCmd.Flags().BoolP("version", "V", false, "Show version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"trace": carapace.ActionValues("0", "1", "2", "3", "4", "5", "6", "7"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
