package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reloadOrRestartCmd = &cobra.Command{
	Use:   "reload-or-restart",
	Short: "Reload one or more units if possible, otherwise start or restart",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reloadOrRestartCmd).Standalone()

	rootCmd.AddCommand(reloadOrRestartCmd)
}
