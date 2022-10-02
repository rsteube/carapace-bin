package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compile_exeCmd = &cobra.Command{
	Use:   "exe",
	Short: "Compile Dart to a self-contained executable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compile_exeCmd).Standalone()

	compile_exeCmd.Flags().StringP("define", "D", "", "Define an environment declaration.")
	compile_exeCmd.Flags().Bool("enable-asserts", false, "Enable assert statements.")
	compile_exeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	compile_exeCmd.Flags().Bool("no-sound-null-safety", false, "Do not respect the nullability of types at runtime.")
	compile_exeCmd.Flags().StringP("output", "o", "", "Write the output to <file name>.")
	compile_exeCmd.Flags().StringP("packages", "p", "", "Get package locations from the specified file instead of .packages.")
	compile_exeCmd.Flags().StringP("save-debugging-info", "S", "", "Remove debugging information from the output and save it separately to the specified file.")
	compile_exeCmd.Flags().Bool("sound-null-safety", false, "Respect the nullability of types at runtime.")
	compile_exeCmd.Flags().String("verbosity", "", "Sets the verbosity level of the compilation.")
	compileCmd.AddCommand(compile_exeCmd)

	carapace.Gen(compile_exeCmd).FlagCompletion(carapace.ActionMap{
		"output":              carapace.ActionFiles(),
		"packages":            carapace.ActionFiles(),
		"save-debugging-info": carapace.ActionFiles(),
		"verbosity": carapace.ActionValuesDescribed(
			"all", "Show all messages",
			"error", "Show only error messages",
			"info", "Show error, warning, and info messages",
			"warning", "Show only error and warning messages",
		),
	})

	carapace.Gen(compile_exeCmd).PositionalCompletion(
		carapace.ActionFiles(".dart"),
	)
}
