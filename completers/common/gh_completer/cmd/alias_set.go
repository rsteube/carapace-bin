package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var alias_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Create a shortcut for a gh command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(alias_setCmd).Standalone()
	alias_setCmd.Flags().BoolP("shell", "s", false, "Declare an alias to be passed through a shell interpreter")
	aliasCmd.AddCommand(alias_setCmd)
}
