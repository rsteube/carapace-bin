package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new or update a GitLab Release for a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_createCmd).Standalone()
	release_createCmd.Flags().StringP("assets-links", "a", "", "`JSON` string representation of assets links (e.g. `--assets='[{\"name\": \"Asset1\", \"url\":\"https://<domain>/some/location/1\", \"link_type\": \"other\", \"filepath\": \"path/to/file\"}]')`")
	release_createCmd.Flags().StringSliceP("milestone", "m", []string{}, "The title of each milestone the release is associated with")
	release_createCmd.Flags().StringP("name", "n", "", "The release name or title")
	release_createCmd.Flags().StringP("notes", "N", "", "The release notes/description. You can use Markdown")
	release_createCmd.Flags().StringP("notes-file", "F", "", "Read release notes `file`. Specify `-` as value to read from stdin")
	release_createCmd.Flags().StringP("ref", "r", "", "If a tag specified doesn't exist, the release is created from ref and tagged with the specified tag name. It can be a commit SHA, another tag name, or a branch name.")
	release_createCmd.Flags().StringP("released-at", "D", "", "The `date` when the release is/was ready. Defaults to the current datetime. Expected in ISO 8601 format (2019-03-15T08:00:00Z)")
	releaseCmd.AddCommand(release_createCmd)

	carapace.Gen(release_createCmd).FlagCompletion(carapace.ActionMap{
		"milestone":  action.ActionMilestones(release_createCmd),
		"notes-file": carapace.ActionFiles(),
		"ref":        action.ActionBranches(release_createCmd), // TODO refs
		// TODO released-at
	})

	carapace.Gen(release_createCmd).PositionalCompletion(
		action.ActionReleases(release_createCmd),
	)

	carapace.Gen(release_createCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
