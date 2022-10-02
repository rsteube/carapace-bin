package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var trust_signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trust_signCmd).Standalone()

	trust_signCmd.Flags().Bool("local", false, "Sign a locally tagged image")
	trustCmd.AddCommand(trust_signCmd)

	carapace.Gen(trust_signCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
