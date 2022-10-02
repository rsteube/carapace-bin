package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pasteBufferCmd = &cobra.Command{
	Use:   "paste-buffer",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pasteBufferCmd).Standalone()

	pasteBufferCmd.Flags().StringS("b", "b", "", "buffer-name")
	pasteBufferCmd.Flags().BoolS("d", "d", false, "TODO description")
	pasteBufferCmd.Flags().BoolS("p", "p", false, "TODO description")
	pasteBufferCmd.Flags().BoolS("r", "r", false, "TODO description")
	pasteBufferCmd.Flags().StringS("s", "s", "", "separator")
	pasteBufferCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(pasteBufferCmd)
}
