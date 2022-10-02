package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var imagetoolsCmd = &cobra.Command{
	Use:   "imagetools",
	Short: "Commands to work on images in registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagetoolsCmd).Standalone()
	rootCmd.AddCommand(imagetoolsCmd)
}
