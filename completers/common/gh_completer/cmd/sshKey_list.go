package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sshKey_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Lists SSH keys in your GitHub account",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_listCmd).Standalone()
	sshKeyCmd.AddCommand(sshKey_listCmd)
}
