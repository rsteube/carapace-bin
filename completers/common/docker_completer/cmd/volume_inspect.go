package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var volume_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Display detailed information on one or more volumes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_inspectCmd).Standalone()

	volume_inspectCmd.Flags().StringP("format", "f", "", "Format the output using the given Go template")
	volumeCmd.AddCommand(volume_inspectCmd)

	carapace.Gen(volume_inspectCmd).PositionalAnyCompletion(docker.ActionVolumes())
}
