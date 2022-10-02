package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var gpgKey_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Lists GPG keys in your GitHub account",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(gpgKey_listCmd).Standalone()
	gpgKeyCmd.AddCommand(gpgKey_listCmd)
}
