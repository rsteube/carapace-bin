package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download asset files from a GitLab Release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_downloadCmd).Standalone()
	release_downloadCmd.Flags().StringArrayP("asset-name", "n", []string{}, "Download only assets that match the name or a glob pattern")
	release_downloadCmd.Flags().StringP("dir", "D", ".", "Directory to download the release assets to")
	release_downloadCmd.Flags().BoolP("include-external", "x", false, "Include external asset files")
	releaseCmd.AddCommand(release_downloadCmd)

	carapace.Gen(release_downloadCmd).FlagCompletion(carapace.ActionMap{
		"asset-name": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionReleaseAssets(release_downloadCmd, c.Args[0]).Invoke(c).Filter(c.Parts).ToA()
			}
			return carapace.ActionValues()
		}),
		"dir": carapace.ActionDirectories(),
	})

	carapace.Gen(release_downloadCmd).PositionalCompletion(
		action.ActionReleases(release_downloadCmd),
	)
}
