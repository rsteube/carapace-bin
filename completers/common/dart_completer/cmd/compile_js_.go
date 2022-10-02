package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compile_jsCmd = &cobra.Command{
	Use:   "js",
	Short: "Compile Dart to JavaScript",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compile_jsCmd).Standalone()

	compile_jsCmd.Flags().StringP("define", "D", "", "Define an environment declaration.")
	compile_jsCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	compile_jsCmd.Flags().BoolP("minified", "m", false, "Generate minified output.")
	compile_jsCmd.Flags().StringP("output", "o", "", "Write the output to <file name>.")
	compile_jsCmd.Flags().String("verbosity", "", "Sets the verbosity level of the compilation.")
	compileCmd.AddCommand(compile_jsCmd)

	carapace.Gen(compile_jsCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
		"verbosity": carapace.ActionValuesDescribed(
			"all", "Show all messages",
			"error", "Show only error messages",
			"info", "Show error, warning, and info messages",
			"warning", "Show only error and warning messages",
		),
	})

	carapace.Gen(compile_jsCmd).PositionalCompletion(
		carapace.ActionFiles(".dart"),
	)
}
