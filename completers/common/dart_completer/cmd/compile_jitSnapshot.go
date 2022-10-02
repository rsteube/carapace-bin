package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compile_jitSnapshotCmd = &cobra.Command{
	Use:   "jit-snapshot",
	Short: "Compile Dart to a JIT snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compile_jitSnapshotCmd).Standalone()

	compile_jitSnapshotCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	compile_jitSnapshotCmd.Flags().StringP("output", "o", "", "Write the output to <file name>.")
	compile_jitSnapshotCmd.Flags().String("verbosity", "", "Sets the verbosity level of the compilation.")
	compileCmd.AddCommand(compile_jitSnapshotCmd)

	carapace.Gen(compile_jitSnapshotCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
		"verbosity": carapace.ActionValuesDescribed(
			"all", "Show all messages",
			"error", "Show only error messages",
			"info", "Show error, warning, and info messages",
			"warning", "Show only error and warning messages",
		),
	})

	carapace.Gen(compile_jitSnapshotCmd).PositionalCompletion(
		carapace.ActionFiles(".dart"),
	)
}
