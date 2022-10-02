package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var volume_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(volume_createCmd).Standalone()

	volume_createCmd.Flags().StringP("driver", "d", "", "Specify volume driver name (default \"local\")")
	volume_createCmd.Flags().String("label", "", "Set metadata for a volume")
	volume_createCmd.Flags().StringP("opt", "o", "", "Set driver specific options (default map[])")
	volumeCmd.AddCommand(volume_createCmd)

	carapace.Gen(volume_createCmd).FlagCompletion(carapace.ActionMap{
		"driver": carapace.ActionValues("local"),
	})

	carapace.Gen(volume_createCmd).PositionalCompletion(docker.ActionVolumes())
}
