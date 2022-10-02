package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var service_psCmd = &cobra.Command{
	Use:   "ps",
	Short: "List the tasks of one or more services",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(service_psCmd).Standalone()

	service_psCmd.Flags().StringP("filter", "f", "", "Filter output based on conditions provided")
	service_psCmd.Flags().String("format", "", "Pretty-print tasks using a Go template")
	service_psCmd.Flags().Bool("no-resolve", false, "Do not map IDs to Names")
	service_psCmd.Flags().Bool("no-trunc", false, "Do not truncate output")
	service_psCmd.Flags().BoolP("quiet", "q", false, "Only display task IDs")
	serviceCmd.AddCommand(service_psCmd)

	carapace.Gen(service_psCmd).PositionalAnyCompletion(docker.ActionServices())
}
