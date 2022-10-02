package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var boxCmd = &cobra.Command{
	Use:   "box",
	Short: "manages boxes: installation, removal, etc.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(boxCmd).Standalone()

	rootCmd.AddCommand(boxCmd)
}
