package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var system_pruneCmd = &cobra.Command{
	Use:   "prune",
	Short: "Remove unused data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_pruneCmd).Standalone()

	system_pruneCmd.Flags().BoolP("all", "a", false, "Remove all unused images not just dangling ones")
	system_pruneCmd.Flags().String("filter", "", "Provide filter values (e.g. 'label=<key>=<value>')")
	system_pruneCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	system_pruneCmd.Flags().Bool("volumes", false, "Prune volumes")
	systemCmd.AddCommand(system_pruneCmd)
}
