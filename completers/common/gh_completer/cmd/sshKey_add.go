package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var sshKey_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an SSH key to your GitHub account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sshKey_addCmd).Standalone()
	sshKey_addCmd.Flags().StringP("title", "t", "", "Title for the new key")
	sshKeyCmd.AddCommand(sshKey_addCmd)

	carapace.Gen(sshKey_addCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
