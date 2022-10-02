package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var extension_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List installed extension commands",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_listCmd).Standalone()
	extensionCmd.AddCommand(extension_listCmd)
}
