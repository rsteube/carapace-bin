package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var box_listCmd = &cobra.Command{
	Use:   "list",
	Short: "list boxes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(box_listCmd).Standalone()

	box_listCmd.Flags().BoolP("box-info", "i", false, "Displays additional information about the boxes")
	boxCmd.AddCommand(box_listCmd)
}
