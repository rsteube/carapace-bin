package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Triggers the agent to reload configuration files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reloadCmd).Standalone()
	addClientFlags(reloadCmd)

	rootCmd.AddCommand(reloadCmd)
}
