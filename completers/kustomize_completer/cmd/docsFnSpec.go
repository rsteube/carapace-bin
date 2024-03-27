package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docsFnSpecCmd = &cobra.Command{
	Use:   "docs-fn-spec",
	Short: "[Alpha] Documentation for Configuration Functions Specification.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsFnSpecCmd).Standalone()
	rootCmd.AddCommand(docsFnSpecCmd)
}
