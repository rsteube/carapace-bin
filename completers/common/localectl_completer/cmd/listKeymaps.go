package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var listKeymapsCmd = &cobra.Command{
	Use:   "list-keymaps",
	Short: "Show known virtual console keyboard mapping",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listKeymapsCmd).Standalone()

	rootCmd.AddCommand(listKeymapsCmd)
}
