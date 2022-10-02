package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var imagetools_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new image based on source images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(imagetools_createCmd).Standalone()
	imagetools_createCmd.Flags().Bool("append", false, "Append to existing manifest")
	imagetools_createCmd.Flags().Bool("dry-run", false, "Show final image instead of pushing")
	imagetools_createCmd.Flags().StringArrayP("file", "f", []string{}, "Read source descriptor from file")
	imagetools_createCmd.Flags().StringArrayP("tag", "t", []string{}, "Set reference for new image")
	imagetoolsCmd.AddCommand(imagetools_createCmd)

	carapace.Gen(imagetools_createCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})

	// TODO verify positional completion - is this correct?
	carapace.Gen(imagetools_createCmd).PositionalAnyCompletion(
		docker.ActionRepositoryTags(),
	)
}
