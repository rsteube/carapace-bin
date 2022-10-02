package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a program",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().Bool("sync", false, "Block until the child process exits")
	rootCmd.AddCommand(execCmd)
}
