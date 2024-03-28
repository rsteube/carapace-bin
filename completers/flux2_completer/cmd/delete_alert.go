package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_alertCmd = &cobra.Command{
	Use:   "alert",
	Short: "Delete a Alert resource",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_alertCmd).Standalone()
	deleteCmd.AddCommand(delete_alertCmd)
}
