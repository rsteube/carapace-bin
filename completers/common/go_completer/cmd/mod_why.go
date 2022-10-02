package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/go_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mod_whyCmd = &cobra.Command{
	Use:   "why",
	Short: "explain why packages or modules are needed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_whyCmd).Standalone()

	mod_whyCmd.Flags().BoolS("m", "m", false, "treat arguments as a list of modules")
	mod_whyCmd.Flags().BoolS("vendor", "vendor", false, "exclude tests of dependencies")
	modCmd.AddCommand(mod_whyCmd)

	carapace.Gen(mod_whyCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				action.ActionPackages(),
				action.ActionModules(true),
			).Invoke(c).Merge().ToA()
		}),
	)
}
