package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/localectl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setLocaleCmd = &cobra.Command{
	Use:   "set-locale",
	Short: "Set system locale",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setLocaleCmd).Standalone()

	rootCmd.AddCommand(setLocaleCmd)

	carapace.Gen(setLocaleCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionLocales().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
