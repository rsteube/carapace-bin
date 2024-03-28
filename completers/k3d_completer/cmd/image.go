package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var imageCmd = &cobra.Command{
	Use:     "image",
	Short:   "Handle container images.",
	Aliases: []string{"images"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imageCmd).Standalone()

	rootCmd.AddCommand(imageCmd)
}
