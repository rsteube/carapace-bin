package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_historyCmd = &cobra.Command{
	Use:   "history",
	Short: "Show the history of an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_historyCmd).Standalone()

	image_historyCmd.Flags().String("format", "", "Pretty-print images using a Go template")
	image_historyCmd.Flags().BoolP("human", "H", false, "Print sizes and dates in human readable format (default true)")
	image_historyCmd.Flags().Bool("no-trunc", false, "Don't truncate output")
	image_historyCmd.Flags().BoolP("quiet", "q", false, "Only show numeric IDs")
	imageCmd.AddCommand(image_historyCmd)

	rootAlias(image_historyCmd, func(cmd *cobra.Command, isAlias bool) {
		carapace.Gen(cmd).PositionalCompletion(
			docker.ActionRepositoryTags(),
		)
	})
}
