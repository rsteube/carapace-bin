package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var service_rollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Revert changes to a service's configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_rollbackCmd).Standalone()

	service_rollbackCmd.Flags().BoolP("detach", "d", false, "Exit immediately instead of waiting for the service to converge")
	service_rollbackCmd.Flags().BoolP("quiet", "q", false, "Suppress progress output")
	serviceCmd.AddCommand(service_rollbackCmd)

	carapace.Gen(service_rollbackCmd).PositionalCompletion(
		docker.ActionServices(),
	)
}
