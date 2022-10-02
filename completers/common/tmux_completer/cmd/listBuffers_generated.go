package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listBuffersCmd = &cobra.Command{
	Use:   "list-buffers",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listBuffersCmd).Standalone()

	listBuffersCmd.Flags().StringS("F", "F", "", "format")
	listBuffersCmd.Flags().StringS("f", "f", "", "filter")
	rootCmd.AddCommand(listBuffersCmd)
}
