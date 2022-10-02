package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var context_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set a property of the current context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_setCmd).Standalone()

	addGlobalOptions(context_setCmd)

	contextCmd.AddCommand(context_setCmd)
}
