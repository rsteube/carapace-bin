package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var trust_revokeCmd = &cobra.Command{
	Use:   "revoke",
	Short: "Remove trust for an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_revokeCmd).Standalone()

	trust_revokeCmd.Flags().BoolP("yes", "y", false, "Do not prompt for confirmation")
	trustCmd.AddCommand(trust_revokeCmd)

	carapace.Gen(trust_revokeCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
