package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var gist_viewCmd = &cobra.Command{
	Use:   "view",
	Short: "View a gist",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gist_viewCmd).Standalone()
	gist_viewCmd.Flags().StringP("filename", "f", "", "Display a single file from the gist")
	gist_viewCmd.Flags().Bool("files", false, "List file names from the gist")
	gist_viewCmd.Flags().BoolP("raw", "r", false, "Print raw instead of rendered gist contents")
	gist_viewCmd.Flags().BoolP("web", "w", false, "Open gist in the browser")
	gistCmd.AddCommand(gist_viewCmd)

	carapace.Gen(gist_viewCmd).FlagCompletion(carapace.ActionMap{
		"filename": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return action.ActionGistFiles(gist_viewCmd, c.Args[0])
			} else {
				return carapace.ActionValues()
			}
		}),
	})

	carapace.Gen(gist_viewCmd).PositionalCompletion(
		action.ActionGists(gist_viewCmd),
	)
}
