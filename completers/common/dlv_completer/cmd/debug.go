package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "Compile and begin debugging main package in current directory, or the package specified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugCmd).Standalone()
	debugCmd.Flags().Bool("continue", false, "Continue the debugged process on start.")
	debugCmd.Flags().String("output", "./__debug_bin", "Output path for the binary.")
	debugCmd.Flags().String("tty", "", "TTY to use for the target program")
	rootCmd.AddCommand(debugCmd)

	// TODO tty
	carapace.Gen(debugCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
