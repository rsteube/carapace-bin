package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload assets to a release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_uploadCmd).Standalone()
	release_uploadCmd.Flags().Bool("clobber", false, "Overwrite existing assets of the same name")
	releaseCmd.AddCommand(release_uploadCmd)

	carapace.Gen(release_uploadCmd).PositionalCompletion(
		action.ActionReleases(release_uploadCmd),
	)
	carapace.Gen(release_uploadCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("#", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionFiles()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
