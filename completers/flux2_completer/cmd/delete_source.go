package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Delete sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_sourceCmd).Standalone()
	deleteCmd.AddCommand(delete_sourceCmd)
}
