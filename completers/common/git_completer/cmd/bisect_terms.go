package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bisect_termsCmd = &cobra.Command{
	Use:   "terms",
	Short: "show the terms used for old and new commits",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bisect_termsCmd).Standalone()
	bisectCmd.Flags().Bool("term-bad", false, "show term for bad/new")
	bisectCmd.Flags().Bool("term-good", false, "show term for good/old")

	bisectCmd.AddCommand(bisect_termsCmd)
}
