package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sshKey_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an SSH key to your GitLab account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_addCmd).Standalone()
	sshKey_addCmd.Flags().StringP("expires-at", "e", "", "The expiration date of the SSH key in ISO 8601 format (YYYY-MM-DDTHH:MM:SSZ)")
	sshKey_addCmd.Flags().StringP("title", "t", "", "New SSH key's title")

	sshKeyCmd.AddCommand(sshKey_addCmd)

	carapace.Gen(sshKey_addCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
