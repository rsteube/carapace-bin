package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/go_completer/cmd/action"
	"github.com/spf13/cobra"
)

var mod_downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download modules to local cache",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mod_downloadCmd).Standalone()

	mod_downloadCmd.Flags().BoolS("json", "json", false, "print a sequence of JSON objects")
	mod_downloadCmd.Flags().BoolS("x", "x", false, "print the commands download executes")
	modCmd.AddCommand(mod_downloadCmd)

	carapace.Gen(mod_downloadCmd).PositionalCompletion(
		action.ActionModules(false),
	)
}
