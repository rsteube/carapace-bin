package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var create_imageCmd = &cobra.Command{
	Use:   "image",
	Short: "Create or update resources dealing with image automation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(create_imageCmd).Standalone()
	createCmd.AddCommand(create_imageCmd)
}
