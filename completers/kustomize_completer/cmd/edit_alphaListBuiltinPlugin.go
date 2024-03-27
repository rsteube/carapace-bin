package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var edit_alphaListBuiltinPluginCmd = &cobra.Command{
	Use:   "alpha-list-builtin-plugin",
	Short: "[Alpha] List the builtin plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(edit_alphaListBuiltinPluginCmd).Standalone()
	editCmd.AddCommand(edit_alphaListBuiltinPluginCmd)
}
