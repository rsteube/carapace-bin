package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var export_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Export image automation objects",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(export_imageCmd).Standalone()
	exportCmd.AddCommand(export_imageCmd)
}
