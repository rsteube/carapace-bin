package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var get_imagesCmd = &cobra.Command{
	Use:   "images",
	Short: "Get image automation object status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(get_imagesCmd).Standalone()
	getCmd.AddCommand(get_imagesCmd)
}
