package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var context_lsCmd = &cobra.Command{
	Use:   "ls",
	Short: "List context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_lsCmd).Standalone()

	context_lsCmd.Flags().String("format", "", "Pretty-print contexts using a Go template")
	context_lsCmd.Flags().BoolP("quiet", "q", false, "Only show context names")
	contextCmd.AddCommand(context_lsCmd)
}
