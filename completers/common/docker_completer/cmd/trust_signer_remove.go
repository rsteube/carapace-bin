package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var trust_signer_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a signer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_signer_removeCmd).Standalone()

	trust_signer_removeCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation before removing the most recent signer")
	trust_signerCmd.AddCommand(trust_signer_removeCmd)

	carapace.Gen(trust_signer_removeCmd).PositionalAnyCompletion(docker.ActionRepositories())
}
