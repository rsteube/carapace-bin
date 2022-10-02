package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var forceLeaveCmd = &cobra.Command{
	Use:   "force-leave",
	Short: "Forces a member of the cluster to enter the left state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forceLeaveCmd).Standalone()
	forceLeaveCmd.Flags().Bool("prune", false, "Remove agent completely from list of members")

	rootCmd.AddCommand(forceLeaveCmd)
}
