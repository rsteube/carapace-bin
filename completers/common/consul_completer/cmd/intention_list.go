package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var intention_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List intentions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(intention_listCmd).Standalone()
	addClientFlags(intention_listCmd)
	addServerFlags(intention_listCmd)
	intention_listCmd.Flags().String("namespace", "", "Specifies the namespace to query")

	intentionCmd.AddCommand(intention_listCmd)

	// TODO namespace completion
}
