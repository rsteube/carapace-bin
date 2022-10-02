package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compile_kernelCmd = &cobra.Command{
	Use:   "kernel",
	Short: "Compile Dart to a kernel snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compile_kernelCmd).Standalone()

	compile_kernelCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	compile_kernelCmd.Flags().StringP("output", "o", "", "Write the output to <file name>.")
	compile_kernelCmd.Flags().String("verbosity", "", "Sets the verbosity level of the compilation.")
	compileCmd.AddCommand(compile_kernelCmd)

	carapace.Gen(compile_kernelCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
		"verbosity": carapace.ActionValuesDescribed(
			"all", "Show all messages",
			"error", "Show only error messages",
			"info", "Show error, warning, and info messages",
			"warning", "Show only error and warning messages",
		),
	})

	carapace.Gen(compile_kernelCmd).PositionalCompletion(
		carapace.ActionFiles(".dart"),
	)
}
