package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compile_aotSnapshotCmd = &cobra.Command{
	Use:   "aot-snapshot",
	Short: "Compile Dart to an AOT snapshot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compile_aotSnapshotCmd).Standalone()

	compile_aotSnapshotCmd.Flags().StringSliceP("define", "D", []string{}, "Define an environment declaration.")
	compile_aotSnapshotCmd.Flags().Bool("enable-asserts", false, "Enable assert statements.")
	compile_aotSnapshotCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	compile_aotSnapshotCmd.Flags().Bool("no-sound-null-safety", false, "Do not respect the nullability of types at runtime.")
	compile_aotSnapshotCmd.Flags().StringP("output", "o", "", "Write the output to <file name>.")
	compile_aotSnapshotCmd.Flags().StringP("packages", "p", "", "Get package locations from the specified file instead of .packages.")
	compile_aotSnapshotCmd.Flags().StringP("save-debugging-info", "S", "", "Remove debugging information from the output and save it separately to the specified file.")
	compile_aotSnapshotCmd.Flags().Bool("sound-null-safety", false, "Respect the nullability of types at runtime.")
	compile_aotSnapshotCmd.Flags().String("verbosity", "", "Sets the verbosity level of the compilation.")
	compileCmd.AddCommand(compile_aotSnapshotCmd)

	carapace.Gen(compile_aotSnapshotCmd).FlagCompletion(carapace.ActionMap{
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

	carapace.Gen(compile_aotSnapshotCmd).PositionalCompletion(
		carapace.ActionFiles(".dart"),
	)
}
