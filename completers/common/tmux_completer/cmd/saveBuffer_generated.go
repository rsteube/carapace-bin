package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var saveBufferCmd = &cobra.Command{
	Use:   "save-buffer",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(saveBufferCmd).Standalone()

	saveBufferCmd.Flags().BoolS("a", "a", false, "TODO description")
	saveBufferCmd.Flags().StringS("b", "b", "", "buffer-name")
	rootCmd.AddCommand(saveBufferCmd)
}
