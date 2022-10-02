package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var imagetools_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Show details of image in the registry",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagetools_inspectCmd).Standalone()
	imagetools_inspectCmd.Flags().Bool("raw", false, "Show original JSON manifest")
	imagetoolsCmd.AddCommand(imagetools_inspectCmd)

	// TODO verify positional completion - is this correct?
	carapace.Gen(imagetools_inspectCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
