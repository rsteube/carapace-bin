package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var suspendCmd = &cobra.Command{
	Use:   "suspend",
	Short: "Suspend resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(suspendCmd).Standalone()
	suspendCmd.PersistentFlags().Bool("all", false, "suspend all resources in that namespace")
	rootCmd.AddCommand(suspendCmd)
}
