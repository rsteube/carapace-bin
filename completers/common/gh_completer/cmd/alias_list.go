package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var alias_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List your aliases",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_listCmd).Standalone()
	aliasCmd.AddCommand(alias_listCmd)
}
