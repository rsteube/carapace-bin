package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Execute a command holding a lock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockCmd).Standalone()
	addClientFlags(lockCmd)
	addServerFlags(lockCmd)

	lockCmd.Flags().Bool("child-exit-code", false, "Exit 2 if the child process exited with an error if this is true.")
	lockCmd.Flags().String("monitor-retry", "", "Number of times to retry if Consul returns a 500 error while monitoring the lock.")
	lockCmd.Flags().String("name", "", "Optional name to associate with the lock session.")
	lockCmd.Flags().Bool("pass-stdin", false, "Pass stdin to the child process.")
	lockCmd.Flags().Bool("shell", false, "Use a shell to run the command.")
	lockCmd.Flags().String("timeout", "", "Maximum amount of time to wait to acquire the lock.")
	lockCmd.Flags().Bool("verbose", false, "Enable verbose (debugging) output.")
	rootCmd.AddCommand(lockCmd)

	// TODO positional completion
}
