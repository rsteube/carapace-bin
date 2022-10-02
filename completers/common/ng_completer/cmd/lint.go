package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Runs linting tools on Angular app code in a given project folder",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lintCmd).Standalone()

	lintCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")

	rootCmd.AddCommand(lintCmd)

	carapace.Gen(lintCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return action.ActionBuilderConfigurations(c.Args[0], "lint").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(lintCmd).PositionalCompletion(
		action.ActionProjects(),
	)
}
