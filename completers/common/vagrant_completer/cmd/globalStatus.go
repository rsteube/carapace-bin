package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var globalStatusCmd = &cobra.Command{
	Use:   "global-status",
	Short: "outputs status Vagrant environments for this user",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalStatusCmd).Standalone()

	globalStatusCmd.Flags().Bool("prune", false, "Prune invalid entries.")
	rootCmd.AddCommand(globalStatusCmd)
}
