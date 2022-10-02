package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var image_loadCmd = &cobra.Command{
	Use:   "load",
	Short: "Load an image from a tar archive or STDIN",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_loadCmd).Standalone()

	image_loadCmd.Flags().StringP("input", "i", "", "Read from tar archive file, instead of STDIN")
	image_loadCmd.Flags().BoolP("quiet", "q", false, "Suppress the load output")
	imageCmd.AddCommand(image_loadCmd)

	rootAlias(image_loadCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
			"input": carapace.ActionFiles(),
		})
	})
}
