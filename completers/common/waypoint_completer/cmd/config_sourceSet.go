package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_sourceSetCmd = &cobra.Command{
	Use:   "source-set",
	Short: "Set the configuration for a dynamic source plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_sourceSetCmd).Standalone()

	config_sourceSetCmd.Flags().String("config", "", "Configuration for the dynamic source type.")
	config_sourceSetCmd.Flags().Bool("delete", false, "Delete the configuration for this source type.")
	config_sourceSetCmd.Flags().String("type", "", "Dynamic source type to configure, such as 'vault'.")

	addGlobalOptions(config_sourceSetCmd)

	configCmd.AddCommand(config_sourceSetCmd)
}
