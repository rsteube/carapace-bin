package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var loadBufferCmd = &cobra.Command{
	Use:   "load-buffer",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(loadBufferCmd).Standalone()

	loadBufferCmd.Flags().StringS("b", "b", "", "buffer-name")
	loadBufferCmd.Flags().StringS("t", "t", "", "target-client")
	rootCmd.AddCommand(loadBufferCmd)
}
