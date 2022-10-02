package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Pack up your CircleCI configuration into a single file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_packCmd).Standalone()
	configCmd.AddCommand(config_packCmd)

	carapace.Gen(config_packCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
