package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var bindKeyCmd = &cobra.Command{
	Use:   "bind-key",
	Short: "bind a key to a command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bindKeyCmd).Standalone()

	bindKeyCmd.Flags().StringS("N", "N", "", "attach a note to the key")
	bindKeyCmd.Flags().StringS("T", "T", "", "specify key table for the binding")
	bindKeyCmd.Flags().BoolS("n", "n", false, "make the binding work without the need for the prefix key")
	bindKeyCmd.Flags().BoolS("r", "r", false, "the key may repeat")
	rootCmd.AddCommand(bindKeyCmd)

	// TODO key table

	carapace.Gen(bindKeyCmd).PositionalCompletion(
		carapace.ActionValues(),
	)

	carapace.Gen(bindKeyCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			c.Args = c.Args[1:]
			return bridge.ActionCarapaceBin("tmux").Invoke(c).ToA()
		}),
	)
}
