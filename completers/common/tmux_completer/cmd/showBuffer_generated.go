package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var showBufferCmd = &cobra.Command{
	Use:   "show-buffer",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showBufferCmd).Standalone()

	showBufferCmd.Flags().StringS("b", "b", "", "buffer-name")
	rootCmd.AddCommand(showBufferCmd)
}
