package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/helm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var get_allCmd = &cobra.Command{
	Use:   "all",
	Short: "download all information for a named release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	get_allCmd.Flags().Int("revision", 0, "get the named release with revision")
	get_allCmd.Flags().String("template", "", "go template for formatting the output, eg: {{.Release.Name}}")
	getCmd.AddCommand(get_allCmd)

	carapace.Gen(get_allCmd).PositionalCompletion(
		action.ActionReleases(),
	)
}
