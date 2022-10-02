package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var image_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display detailed information on one or more images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(image_inspectCmd).Standalone()

	image_inspectCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	imageCmd.AddCommand(image_inspectCmd)

	carapace.Gen(image_inspectCmd).PositionalAnyCompletion(docker.ActionRepositoryTags())
}
