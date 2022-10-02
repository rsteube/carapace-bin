package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var intention_matchCmd = &cobra.Command{
	Use:   "match",
	Short: "Show intentions that match a source or destination",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(intention_matchCmd).Standalone()

	intention_matchCmd.Flags().Bool("destination", false, "Match intentions with the given destination.")
	intention_matchCmd.Flags().String("namespace", "", "Specifies the namespace to query.")
	intention_matchCmd.Flags().Bool("source", false, "Match intentions with the given source.")
	intentionCmd.AddCommand(intention_matchCmd)

	// TODO flag completion

	// TODO  positional completion
}
