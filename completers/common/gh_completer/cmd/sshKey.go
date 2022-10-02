package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sshKeyCmd = &cobra.Command{
	Use:   "ssh-key",
	Short: "Manage SSH keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKeyCmd).Standalone()
	rootCmd.AddCommand(sshKeyCmd)
}
