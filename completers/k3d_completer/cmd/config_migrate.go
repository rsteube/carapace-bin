package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_migrateCmd = &cobra.Command{
	Use:     "migrate INPUT [OUTPUT]",
	Short:   "",
	Aliases: []string{"update"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_migrateCmd).Standalone()

	configCmd.AddCommand(config_migrateCmd)
}
