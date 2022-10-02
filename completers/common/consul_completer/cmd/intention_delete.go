package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var intention_deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an intention",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(intention_deleteCmd).Standalone()
	addClientFlags(intention_deleteCmd)
	addServerFlags(intention_deleteCmd)
	intention_deleteCmd.Flags().String("namespace", "", "Specifies the namespace to query")

	intentionCmd.AddCommand(intention_deleteCmd)

	// TODO namespace completion

	// TODO positional completion
}
