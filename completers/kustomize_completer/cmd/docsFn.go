package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var docsFnCmd = &cobra.Command{
	Use:   "docs-fn",
	Short: "[Alpha] Documentation for developing and invoking Configuration Functions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(docsFnCmd).Standalone()
	rootCmd.AddCommand(docsFnCmd)
}
