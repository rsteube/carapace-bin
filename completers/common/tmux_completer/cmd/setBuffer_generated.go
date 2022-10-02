package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setBufferCmd = &cobra.Command{
	Use:   "set-buffer",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setBufferCmd).Standalone()

	setBufferCmd.Flags().BoolS("a", "a", false, "TODO description")
	setBufferCmd.Flags().StringS("b", "b", "", "buffer-name")
	setBufferCmd.Flags().StringS("n", "n", "", "new-buffer-name")
	setBufferCmd.Flags().StringS("t", "t", "", "target-client")
	setBufferCmd.Flags().BoolS("w", "w", false, "TODO description")
	rootCmd.AddCommand(setBufferCmd)
}
