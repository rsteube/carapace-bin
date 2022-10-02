package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var context_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List contexts",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_listCmd).Standalone()

	addGlobalOptions(context_listCmd)

	contextCmd.AddCommand(context_listCmd)
}
