package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var delete_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Delete image automation objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(delete_imageCmd).Standalone()
	deleteCmd.AddCommand(delete_imageCmd)
}
