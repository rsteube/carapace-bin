package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "Execute a command in the context of a running application instance",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().String("instance", "", "Start an exec session on this specific instance")
	addGlobalOptions(execCmd)
	rootCmd.AddCommand(execCmd)
}
