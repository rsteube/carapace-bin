package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var thawCmd = &cobra.Command{
	Use:   "thaw",
	Short: "Resume execution of a frozen unit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(thawCmd).Standalone()

	rootCmd.AddCommand(thawCmd)

	carapace.Gen(thawCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionUnits().Invoke(c).Filter(c.Args).ToA()
		}),
	)
}
