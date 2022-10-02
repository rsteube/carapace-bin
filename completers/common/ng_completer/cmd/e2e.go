package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/ng_completer/cmd/action"
	"github.com/spf13/cobra"
)

var e2eCmd = &cobra.Command{
	Use:   "e2e",
	Short: "Builds and serves an Angular app, then runs end-to-end tests",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(e2eCmd).Standalone()

	e2eCmd.Flags().StringP("configuration", "c", "", "One or more named builder configurations")
	rootCmd.AddCommand(e2eCmd)

	carapace.Gen(e2eCmd).FlagCompletion(carapace.ActionMap{
		"configuration": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			if len(c.Args) < 1 {
				return carapace.ActionValues()
			}
			return action.ActionBuilderConfigurations(c.Args[0], "").Invoke(c).Filter(c.Parts).ToA()
		}),
	})

	carapace.Gen(e2eCmd).PositionalCompletion(
		action.ActionProjects(),
	)
}
