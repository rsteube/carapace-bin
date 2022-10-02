package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View information about a release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_viewCmd).Standalone()
	release_viewCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	release_viewCmd.Flags().StringSlice("json", []string{}, "Output JSON with the specified `fields`")
	release_viewCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template")
	release_viewCmd.Flags().BoolP("web", "w", false, "Open the release in the browser")
	releaseCmd.AddCommand(release_viewCmd)

	carapace.Gen(release_viewCmd).FlagCompletion(carapace.ActionMap{
		"json": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return action.ActionReleaseFields().Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(release_viewCmd).PositionalCompletion(
		action.ActionReleases(release_viewCmd),
	)
}
