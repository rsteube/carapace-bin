package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a Dart program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().StringP("define", "D", "", "Define an environment declaration.")
	runCmd.Flags().Bool("enable-asserts", false, "Enable assert statements.")
	runCmd.Flags().String("enable-vm-service", "", "Enables the VM service and listens on the specified port for connections")
	runCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	runCmd.Flags().Bool("no-enable-asserts", false, "Enable assert statements.")
	runCmd.Flags().Bool("no-pause-isolates-on-exit", false, "Do not pause isolates on exit when running with --enable-vm-service.")
	runCmd.Flags().Bool("no-pause-isolates-on-start", false, "Pause isolates on start when running with --enable-vm-service.")
	runCmd.Flags().Bool("no-pause-isolates-on-unhandled-exceptions", false, "Do not pause isolates when an unhandled exception is encountered when running with --enable-vm-service.")
	runCmd.Flags().Bool("no-warn-on-pause-with-no-debugger", false, "Do not print a warning when an isolate pauses with no attached debugger when running with --enable-vm-service.")
	runCmd.Flags().String("observe", "", "The observe flag is a convenience flag used to run a program with a set of common options useful for debugging.")
	runCmd.Flags().Bool("pause-isolates-on-exit", false, "Pause isolates on exit when running with --enable-vm-service.")
	runCmd.Flags().Bool("pause-isolates-on-start", false, "Pause isolates on start when running with --enable-vm-service.")
	runCmd.Flags().Bool("pause-isolates-on-unhandled-exceptions", false, "Pause isolates when an unhandled exception is encountered when running with --enable-vm-service.")
	runCmd.Flags().String("verbosity", "", "Sets the verbosity level of the compilation.")
	runCmd.Flags().Bool("warn-on-pause-with-no-debugger", false, "Print a warning when an isolate pauses with no attached debugger when running with --enable-vm-service.")
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"verbosity": carapace.ActionValuesDescribed(
			"all", "Show all messages",
			"error", "Show only error messages",
			"info", "Show error, warning, and info messages",
			"warning", "Show only error and warning messages",
		),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionFiles(".dart"),
	)
}
