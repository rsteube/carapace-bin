package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sshKey_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists currently authenticated user’s SSH keys",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_listCmd).Standalone()
	sshKey_listCmd.Flags().Bool("show-id", false, "Show IDs of SSH Keys")
	sshKeyCmd.AddCommand(sshKey_listCmd)
}
