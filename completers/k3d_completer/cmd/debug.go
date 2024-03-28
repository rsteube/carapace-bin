package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var debugCmd = &cobra.Command{
	Use:    "debug",
	Short:  "Debug k3d cluster(s)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugCmd).Standalone()

	rootCmd.AddCommand(debugCmd)
}
