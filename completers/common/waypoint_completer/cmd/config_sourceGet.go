package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var config_sourceGetCmd = &cobra.Command{
	Use:   "source-get",
	Short: "Get the configuration for a dynamic source plugin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_sourceGetCmd).Standalone()

	config_sourceGetCmd.Flags().String("type", "", "Dynamic source type to look up, such as 'vault'.")

	addGlobalOptions(config_sourceGetCmd)

	configCmd.AddCommand(config_sourceGetCmd)
}
