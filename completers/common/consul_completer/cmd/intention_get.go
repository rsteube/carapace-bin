package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var intention_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Show information about an intention.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(intention_getCmd).Standalone()
	addClientFlags(intention_getCmd)
	addServerFlags(intention_getCmd)
	intention_getCmd.Flags().String("namespace", "", "Specifies the namespace to query")

	intentionCmd.AddCommand(intention_getCmd)

	// TODO namespace completion

	// TODO positional completion
}
