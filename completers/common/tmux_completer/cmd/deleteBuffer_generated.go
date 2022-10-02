package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var deleteBufferCmd = &cobra.Command{
	Use:   "delete-buffer",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteBufferCmd).Standalone()

	deleteBufferCmd.Flags().StringS("b", "b", "", "buffer-name")
	rootCmd.AddCommand(deleteBufferCmd)
}
