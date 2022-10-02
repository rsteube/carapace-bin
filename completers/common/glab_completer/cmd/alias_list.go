package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var alias_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the available aliases.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_listCmd).Standalone()
	aliasCmd.AddCommand(alias_listCmd)
}
