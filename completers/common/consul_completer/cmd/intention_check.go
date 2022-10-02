package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var intention_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check whether a connection between two services is allowed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(intention_checkCmd).Standalone()
	addClientFlags(intention_checkCmd)
	addServerFlags(intention_checkCmd)
	intention_checkCmd.Flags().String("namespace", "", "Specifies the namespace to query")

	intentionCmd.AddCommand(intention_checkCmd)

	// TODO namespace completion

	// TODO positional completion
}
