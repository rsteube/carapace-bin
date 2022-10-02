package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_uploadCmd = &cobra.Command{
	Use:   "upload",
	Short: "Upload release asset files or links to GitLab Release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_uploadCmd).Standalone()
	release_uploadCmd.Flags().StringP("assets-links", "a", "", "`JSON` string representation of assets links (e.g. `--assets='[{\"name\": \"Asset1\", \"url\":\"https://<domain>/some/location/1\", \"link_type\": \"other\", \"filepath\": \"path/to/file\"}]')`")
	releaseCmd.AddCommand(release_uploadCmd)

	carapace.Gen(release_uploadCmd).PositionalCompletion(
		action.ActionReleases(release_uploadCmd),
	)

	carapace.Gen(release_uploadCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
