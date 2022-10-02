package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var assembleCmd = &cobra.Command{
	Use:   "assemble",
	Short: "Assemble and build Flutter resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(assembleCmd).Standalone()

	assembleCmd.Flags().StringArray("DartDefines", []string{}, "Additional key-value pairs that will be available as constants.")
	assembleCmd.Flags().String("build-inputs", "", "A file path where a newline-separated file containing all inputs used will be written after a build.")
	assembleCmd.Flags().String("build-outputs", "", "A file path where a newline-separated file containing all outputs created will be written after a build.")
	assembleCmd.Flags().StringP("define", "d", "", "Allows passing configuration to a target, as in \"--define=target=key=value\".")
	assembleCmd.Flags().String("depfile", "", "A file path where a depfile will be written.")
	assembleCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	assembleCmd.Flags().StringP("input", "i", "", "Allows passing additional inputs with \"--input=key=value\".")
	assembleCmd.Flags().StringP("output", "o", "", "A directory where output files will be written.")
	assembleCmd.Flags().String("performance-measurement-file", "", "Output individual target performance to a JSON file.")
	assembleCmd.Flags().String("resource-pool-size", "", "The maximum number of concurrent tasks the build system will run.")
	rootCmd.AddCommand(assembleCmd)

	carapace.Gen(assembleCmd).FlagCompletion(carapace.ActionMap{
		"build-inputs":                 carapace.ActionFiles(),
		"build-outputs":                carapace.ActionFiles(),
		"depfile":                      carapace.ActionFiles(),
		"output":                       carapace.ActionDirectories(),
		"performance-measurement-file": carapace.ActionFiles(),
	})

	carapace.Gen(assembleCmd).PositionalCompletion(
		action.ActionTargets(),
	)
}
