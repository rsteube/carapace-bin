package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a precompiled binary, and begin a debug session.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()
	execCmd.Flags().Bool("continue", false, "Continue the debugged process on start.")
	execCmd.Flags().String("tty", "", "TTY to use for the target program")
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
