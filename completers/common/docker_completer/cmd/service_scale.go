package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var service_scaleCmd = &cobra.Command{
	Use:   "scale",
	Short: "Scale one or multiple replicated services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_scaleCmd).Standalone()

	service_scaleCmd.Flags().BoolP("detach", "d", false, "Exit immediately instead of waiting for the service to converge")
	serviceCmd.AddCommand(service_scaleCmd)

	carapace.Gen(service_scaleCmd).PositionalAnyCompletion(docker.ActionServices())
}
