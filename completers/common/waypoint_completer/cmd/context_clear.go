package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var context_clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Unset the default context",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_clearCmd).Standalone()

	addGlobalOptions(context_clearCmd)

	contextCmd.AddCommand(context_clearCmd)
}
